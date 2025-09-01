package functions

import (
	"bufio"
	"strings"
	"time"

	gb "netCat/global"
)

func ClientReader(client *gb.Client, reader *bufio.Reader, leave chan<- *gb.Client, messages chan<- *gb.Message) {
	for {
		line, err := reader.ReadString('\n')
		if err != nil {
			leave <- client
			return
		}
		text := strings.TrimRight(line, "\r\n")
		if strings.TrimSpace(text) == "" {
			continue
		}
		messages <- &gb.Message{From: client.Name, Text: text, Timestamp: time.Now()}

	}
}
