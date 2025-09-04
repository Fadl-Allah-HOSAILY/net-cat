package functions

import (
	"strings"
	"time"

	gb "netCat/global"
)

func ClientReader(client gb.Client) {
	defer client.Conn.Close()
	for {
		readTime := time.Now().Format("2006-01-02 15:04:05")
		format := "[" + readTime + "]" + "[" + client.Name + "]:"
		_, err := client.Conn.Write([]byte(format))
		if err != nil {
			return
		}
		buf := make([]byte, 1024)
		_, err = client.Conn.Read(buf)
		if err != nil {
			CloseConnection(client, clients)
			return
		}
		text := (strings.Split(string(buf), "\n"))

		if text[0] == "" {
			continue
		}
		if !IsLetter(text[0]){
			client.Conn.Write([]byte("Invalide message \n"))
			buf=[]byte{}
			continue
		}
		writeTime := time.Now().Format("2006-01-02 15:04:05")
		textFormat := "[" + writeTime + "]" + "[" + client.Name + "]:" + text[0]
		history = append(history, textFormat)
		for name, clientConn := range clients {
			format := "[" + writeTime + "]" + "[" + name + "]:"
			if name != client.Name {
				clientConn.Write([]byte("\n" + textFormat))
				_, err := clientConn.Write([]byte(format))
				if err != nil {
					return
				}
			}

		}
	}
}
