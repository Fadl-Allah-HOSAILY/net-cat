package functions

import (
	"bufio"
	"net"
	"strings"
)

func GetClientName(conn net.Conn, reader *bufio.Reader) string {
	for {
		line, err := reader.ReadString('\n')
		if err != nil {
			return ""
		}
		name := strings.TrimSpace(line)
		if name != "" {
			return name
		}
		conn.Write([]byte("Name cannot be empty. \n[ENTER YOUR NAME]: "))
	}
}
