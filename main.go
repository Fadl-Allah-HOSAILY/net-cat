package main

import (
	"fmt"

	fn "netCat/functions"
)

const DefaultPort = "8989"

func main() {
	port := fn.ParsePort(DefaultPort)
	ln := fn.StartServer(port)
	defer ln.Close()

	fmt.Println("Listening on the port:", port)

	for {
		conn, err := ln.Accept()
		if err != nil {
			fmt.Println("accept error:", err)
			continue
		}
		go fn.HandleConnection(conn)
	}
}
