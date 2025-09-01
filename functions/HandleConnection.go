package functions

import (
	"bufio"
	"net"

	gb "netCat/global"
)

var linuxLogo = `Welcome to TCP-Chat!
	 _nnnn_
	dGGGGMMb
   @p~qp~~qMb
   M|@||@) M|
   @,----.JM|
  JS^\__/  qKL
 dZP        qKRb
dZP          qKKb
fZP            SMMb
HZM            MMMM
FqM            MMMM
__| ".        |\dS"qML
|    '.       | '' \Zq
_)      \.___.,|     .'
\____   )MMMMMP|   .'
 '-'       '--'
[ENTER YOUR NAME]: `

func HandleConnection(conn net.Conn, join chan<- *gb.Client, leave chan<- *gb.Client, messages chan<- *gb.Message) {
	defer conn.Close()

	conn.Write([]byte(linuxLogo))

	reader := bufio.NewReader(conn)
	name := GetClientName(conn, reader)
	if name == "" {
		return
	}

	client := &gb.Client{Conn: conn, Name: name, Ch: make(chan string, 32)}
	go ClientWriter(client)

	join <- client
	ClientReader(client, reader, leave, messages)
}
