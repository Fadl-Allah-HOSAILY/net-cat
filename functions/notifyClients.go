package functions

import (
	"net"

	gb "netCat/global"
)

func NotifyClients(clients map[net.Conn]*gb.Client, msg string, exclude net.Conn) {
	for _, cl := range clients {
		if exclude != nil && cl.Conn == exclude {
			continue
		}
		cl.Ch <- msg
	}
}
