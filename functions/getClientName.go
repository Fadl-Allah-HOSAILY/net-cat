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

		if name == "" {
			conn.Write([]byte("Name cannot be empty. \n[ENTER YOUR NAME]: "))
			continue
		}
		if !IsLetter(name) {
			conn.Write([]byte("Invalide ame. \n[ENTER YOUR NAME]: "))
			continue
		}

		// Vérifier si le nom existe déjà
		mu.Lock()
		_, exists := clients[name]
		mu.Unlock()

		if exists {
			conn.Write([]byte("Name already exist. \n[ENTER YOUR NAME]: "))
			continue
		}

		return name
	}
}
