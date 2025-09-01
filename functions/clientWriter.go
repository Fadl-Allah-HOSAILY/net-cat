package functions

import gb "netCat/global"

func ClientWriter(client *gb.Client) {
	for msg := range client.Ch {
		_, err := client.Conn.Write([]byte(msg + "\n"))
		if err != nil {
			return
		}
	}
}
