package main

import (
	"fmt"
	"os"
	"sync"

	fn "netCat/functions"
	gb "netCat/global"
)

const DefaultPort = "8989"

func main() {
	port := fn.ParsePort(DefaultPort)
	ln := fn.StartServer(port)
	defer ln.Close()

	fmt.Println("Listening on the port:", port)

	join := make(chan *gb.Client)
	leave := make(chan *gb.Client)
	message := make(chan *gb.Message)

	history := make([]string, 0)
	var historyMu sync.Mutex

	go fn.Broadcaster(join, leave, message, &history, &historyMu)

	for {
		conn, err := ln.Accept()
		if err != nil {
			fmt.Fprintf(os.Stderr, "accept error: %v\n", err)
			fmt.Println("accept error:", err)
			continue
		}
		go fn.HandleConnection(conn, join, leave, message)
	}
}
