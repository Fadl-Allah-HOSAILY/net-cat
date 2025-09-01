package functions

import (
	"net"

	gb "netCat/global"
)

func Broadcast(clients map[net.Conn]*gb.Client, line string, sender *gb.Client) {
	for _, cl := range clients {
		if cl == sender {
			continue
		}
		cl.Ch <- line
	}
}
