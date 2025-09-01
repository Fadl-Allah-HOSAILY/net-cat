package functions

import (
	"fmt"
	"net"
	"os"
)

func StartServer(port string) net.Listener {
	ln, err := net.Listen("tcp", ":"+port)
	if err != nil {
		fmt.Println("failed to listen on port:", port)
		os.Exit(1)
	}
	return ln
}
