package server

import (
	"bufio"
	"encoding/json"
	"fmt"
	"net"

	"github.com/RodrigoGonzalez78/sockets/models"
)

var clients []models.Client

func StartServer(direccion string) {

	ln, err := net.Listen("tcp", direccion)

	if err != nil {
		fmt.Println("Error al crear el servidor:", err)
		return
	}
	defer ln.Close()

	fmt.Println("Servidor escuchando en la dirección:", direccion)

	for {
		conn, err := ln.Accept()

		if err != nil {
			fmt.Println("Error al aceptar la conexión:", err)
			continue
		}

		client := models.Client{
			Connection: conn,
		}

		clients = append(clients, client)

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
		var mensaje models.Mensaje

		err := json.NewDecoder(reader).Decode(&mensaje)

		if err != nil {
			fmt.Println("Error al leer el mensaje del cliente:", err)
			return
		}

		fmt.Printf("\n## %s : %s  %v:%v ##\n", mensaje.NombreCliente, mensaje.Mensaje, mensaje.FechaHora.Hour(), mensaje.FechaHora.Minute())

		// Enviar el mensaje a todos los clientes
		broadcastMessage(mensaje, conn.RemoteAddr())
	}
}

func broadcastMessage(mensaje models.Mensaje, direccion net.Addr) {

	for _, client := range clients {
		if client.Connection.RemoteAddr() != direccion {
			json.NewEncoder(client.Connection).Encode(mensaje)
		}

	}
}
