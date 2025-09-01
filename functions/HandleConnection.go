package functions

import (
	"bufio"
	"net"
	"sync"

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

func HandleConnection(
	conn net.Conn,
	join chan<- *gb.Client,
	leave chan<- *gb.Client,
	messages chan<- *gb.Message,
	existingNames map[string]bool,
	namesMu *sync.Mutex,
) {
	defer conn.Close()

	conn.Write([]byte(linuxLogo))

	reader := bufio.NewReader(conn)
	name := GetClientName(conn, reader, existingNames, namesMu)
	if name == "" {
		return
	}

	client := &gb.Client{Conn: conn, Name: name, Ch: make(chan string, 32)}
	go ClientWriter(client)

	// Ajouter le nom à la liste des noms existants
	namesMu.Lock()
	existingNames[name] = true
	namesMu.Unlock()

	join <- client

	// Lecture des messages
	ClientReader(client, reader, leave, messages)

	// Quand le client part, supprimer son nom
	namesMu.Lock()
	delete(existingNames, name)
	namesMu.Unlock()
}
