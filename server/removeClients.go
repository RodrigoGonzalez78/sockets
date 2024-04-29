package server

import (
	"fmt"
	"net"
)

func removeClient(conn net.Conn) {

	client, ok := clients[conn.RemoteAddr().String()]

	if !ok {
		return
	}

	delete(clients, conn.RemoteAddr().String())
	fmt.Println("Cliente desconectado:", conn.RemoteAddr())

	client.Connection.Close()

}
