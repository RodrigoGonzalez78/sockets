package server

import (
	"bufio"
	"encoding/json"
	"fmt"
	"net"

	"github.com/RodrigoGonzalez78/sockets_messages/models"
)

var clients map[string]models.Client

func StartServer(address string) {

	clients = make(map[string]models.Client)

	ln, err := net.Listen("tcp", address)

	if err != nil {
		fmt.Println("Error al crear el servidor:", err)
		return
	}
	defer ln.Close()

	fmt.Println("Servidor escuchando en la dirección:", address)

	for {
		conn, err := ln.Accept()
		if err != nil {
			fmt.Println("Error al aceptar la conexión:", err)
			continue
		}

		client := models.Client{
			Connection: conn,
		}

		clients[conn.RemoteAddr().String()] = client

		go handleConnection(conn)
	}
}

func handleConnection(conn net.Conn) {

	defer conn.Close()

	fmt.Println("Nueva conexión establecida con el host:", conn.RemoteAddr())

	reader := bufio.NewReader(conn)

	for {
		var transmition models.Transmition

		err := json.NewDecoder(reader).Decode(&transmition)

		if err != nil {
			fmt.Println("Error al leer el mensaje del cliente:", err)
			removeClient(conn)
			return
		}

		if transmition.Operation == "send_message" {
			sendMessageToClient(transmition, conn)
		} else if transmition.Operation == "disconect" {
			removeClient(conn)
		} else if transmition.Operation == "get_clients_list" {

		}

	}
}

func removeClient(conn net.Conn) {

	delete(clients, conn.RemoteAddr().String())
	fmt.Println("Cliente desconectado:", conn.RemoteAddr())

}

func sendMessageToClient(transmition models.Transmition, sender net.Conn) {

	client, ok := clients[transmition.Messaje.IP]

	if !ok {
		fmt.Println("El cliente destino no está conectado:", transmition.Messaje.IP)
		return
	}

	transmition.Messaje.IP = sender.RemoteAddr().String()
	transmition.Operation = "receive_message"

	json.NewEncoder(client.Connection).Encode(transmition.Messaje.IP)
}
