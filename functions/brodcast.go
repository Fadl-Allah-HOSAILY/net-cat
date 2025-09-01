package functions

import (
	"net"

	gb "netCat/global"
)

func Broadcast(clients map[net.Conn]*gb.Client, line string) {
	for _, cl := range clients {
		cl.Ch <- line
	}
}
