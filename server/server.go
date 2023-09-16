package server

import (
	"bufio"
	"fmt"
	"net"
)

func StartServer() {
	ln, err := net.Listen("tcp", ":8080")
	if err != nil {
		fmt.Println("Error al crear el servidor:", err)
		return
	}
	defer ln.Close()

	fmt.Println("Servidor escuchando en el puerto 8080")

	for {
		conn, err := ln.Accept()
		if err != nil {
			fmt.Println("Error al aceptar la conexión:", err)
			continue
		}

		go handleConnection(conn)
	}
}

func handleConnection(conn net.Conn) {
	defer conn.Close()

	fmt.Println("Nueva conexión establecida:", conn.RemoteAddr())

	// Crear un lector para leer mensajes del cliente
	reader := bufio.NewReader(conn)

	for {
		// Leer el mensaje del cliente
		message, err := reader.ReadString('\n')
		if err != nil {
			fmt.Println("Error al leer el mensaje del cliente:", err)
			return
		}

		fmt.Println("Mensaje del cliente:", message)

		// Enviar una respuesta al cliente
		conn.Write([]byte("Mensaje recibido: " + message))
	}
}
