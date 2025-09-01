package functions

import (
	"fmt"
	"net"
	"sync"

	gb "netCat/global"
)

func Broadcaster(join <-chan *gb.Client, leave <-chan *gb.Client, messages <-chan *gb.Message, history *[]string, historyMu *sync.Mutex) {
	clients := make(map[net.Conn]*gb.Client)
	const maxHistory = 100

	for {
		select {
		case c := <-join:
			if len(clients) >= 10 {
				c.Conn.Write([]byte("Server full. Max connections reached.\n"))
				c.Conn.Close()
				continue
			}
			clients[c.Conn] = c

			SendHistory(c, history, historyMu)
			NotifyClients(clients, fmt.Sprintf("%s has joined our chat...", c.Name), c.Conn)

		case c := <-leave:
			delete(clients, c.Conn)
			NotifyClients(clients, fmt.Sprintf("%s has left our chat...", c.Name), nil)
			close(c.Ch)

		case m := <-messages:
			line := FormatMessage(m)
			AppendHistory(line, history, historyMu, maxHistory)
			var sender *gb.Client
			for _, cl := range clients {
				if cl.Name == m.From {
					sender = cl
					break
				}
			}

			Broadcast(clients, line, sender)
		}
	}
}
