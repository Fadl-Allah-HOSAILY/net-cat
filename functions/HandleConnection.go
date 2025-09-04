package functions

import (
	"bufio"
	"net"
	"sync"

	gb "netCat/global"
)

var mu sync.Mutex

func HandleConnection(conn net.Conn) {
	linuxLogo := `
Welcome to TCP-Chat!
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
 |    .        ' \Zq
_)      \.___.,|     .'
\____   )MMMMMP|   .'
     -'       '
[ENTER YOUR NAME]: `

	defer conn.Close()

	conn.Write([]byte(linuxLogo))

	reader := bufio.NewReader(conn)
	name := GetClientName(conn, reader)
	client := gb.Client{Conn: conn, Name: name}
	if len(clients) >= 10 {
		client.Conn.Write([]byte("The chat is full"))
		return
	}

	mu.Lock()

	clients[client.Name] = client.Conn
	mu.Unlock()
	OpenConnection(client, clients)
	for _, message := range history {
		conn.Write([]byte(message))
	}
	ClientReader(client)
}
