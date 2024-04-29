package server

import (
	"encoding/json"
	"fmt"
	"net"

	"github.com/RodrigoGonzalez78/sockets_messages/models"
)

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
