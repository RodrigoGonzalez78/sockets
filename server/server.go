package server

import (
	"bufio"
	"encoding/json"
	"fmt"
	"net"
	"time"

	"github.com/RodrigoGonzalez78/sockets_messages/models"
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

	fmt.Println("Nueva conexión establecida con el host:", conn.RemoteAddr())

	reader := bufio.NewReader(conn)

	for {
		var mensaje models.Message

		err := json.NewDecoder(reader).Decode(&mensaje)

		if err != nil {
			fmt.Println("Error al leer el mensaje del cliente:", err)

			removeClient(conn)
			return
		}

		fmt.Printf("\n## %s : %s %v ##\n", mensaje.ClientName, mensaje.Message, mensaje.Time)

		if mensaje.Message == "/listar" {
			sendConnectedClientsList(conn)
		} else if mensaje.Message == "/quitar" {
			sendDisconnectMessage(conn)
			return
		} else {
			broadcastMessage(mensaje, conn.RemoteAddr())
		}
	}
}

func sendConnectedClientsList(conn net.Conn) {

	clientesString := "Clientes conectados:\n"
	for _, client := range clients {
		clientesString += "- " + client.Connection.RemoteAddr().String() + "\n"
	}

	listaClientes := models.Message{
		ClientName: "Servidor",
		Message:    clientesString,
		Time:       time.Now().Format("15:04"),
	}
	json.NewEncoder(conn).Encode(listaClientes)

}

func removeClient(conn net.Conn) {
	// Encuentra y elimina al cliente de la lista de clientes
	for i, client := range clients {

		if client.Connection.RemoteAddr() == conn.RemoteAddr() {
			// Elimina al cliente de la lista
			clients = append(clients[:i], clients[i+1:]...)
			fmt.Println("Cliente desconectado: ", conn.RemoteAddr())
			break
		}
	}

}

func sendDisconnectMessage(conn net.Conn) {
	closeConexion := models.Message{
		ClientName: "Servidor",
		Message:    "Tu sesión se ha cerrado.",
		Time:       time.Now().Format("15:04"),
	}
	json.NewEncoder(conn).Encode(closeConexion)
	removeClient(conn)
	conn.Close()
}

func broadcastMessage(message models.Message, direction net.Addr) {

	for _, client := range clients {
		if client.Connection.RemoteAddr() != direction {

			message.Time = time.Now().Format("15:04")

			json.NewEncoder(client.Connection).Encode(message)
		}

	}
}
