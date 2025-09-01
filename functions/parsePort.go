package functions

import (
	"fmt"
	"os"
)

func ParsePort(DefaultPort string) string {
	port := DefaultPort
	args := os.Args[1:]
	if len(args) == 1 {
		port = args[0]
	} else if len(args) > 1 {
		fmt.Println("[USAGE]: ./TCPChat $port")
		os.Exit(1)
	}
	return port
}
