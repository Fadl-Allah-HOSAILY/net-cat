package main

import (
	"fmt"

	fn "netCat/functions"
	"netCat/utils"
)

const DefaultPort = "8989"

func main() {
	port := utils.ParsePort(DefaultPort)
	ln := utils.StartServer(port)
	defer ln.Close()

	fmt.Println("Listening on the port:", port)

	existingNames := make(map[string]bool)

	for {
		conn, err := ln.Accept()
		if err != nil {
			fmt.Println("accept error:", err)
			continue
		}
		go fn.HandleConnection(conn, existingNames)
	}
}
