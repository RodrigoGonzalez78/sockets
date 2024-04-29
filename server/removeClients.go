package server

import (
	"fmt"
	"net"
)

func removeClient(conn net.Conn) {

	delete(clients, conn.RemoteAddr().String())
	fmt.Println("Cliente desconectado:", conn.RemoteAddr())

}
