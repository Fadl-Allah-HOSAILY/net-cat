package global

import (
	"net"
	"time"
)

// Client represents a connected client
type Client struct {
	Conn net.Conn
	Name string
	Ch   chan string
}

// Message represents a chat message
type Message struct {
	From      string
	Text      string
	Timestamp time.Time
}
