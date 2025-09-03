package functions

import (
	"net"
	"time"

	gb "netCat/global"
)

var (
	history []string
	clients = make(map[string]net.Conn)
)

func OpenConnection(client gb.Client, clients map[string]net.Conn) {
	for name, cl := range clients {
		if name != client.Name {
			cl.Write([]byte("\n" + client.Name + " has joined our chat..."))

			writeTime := time.Now().Format("2006-01-02 15:04:05")
			format := "[" + writeTime + "]" + "[" + name + "]:"
			_, err := cl.Write([]byte("\n" + format))
			if err != nil {
				continue
			}
		}
	}
}

func CloseConnection(client gb.Client, clients map[string]net.Conn) {
	for name, cl := range clients {
		if name != client.Name {
			cl.Write([]byte("\n" + client.Name + "has left our chat..."))
			writeTime := time.Now().Format("2006-01-02 15:04:05")
			format := "[" + writeTime + "]" + "[" + name + "]:"
			_, err := cl.Write([]byte("\n" + format))
			if err != nil {
				continue
			}
		}
	}
	mu.Lock()
	delete(clients, client.Name)
	mu.Unlock()
}
