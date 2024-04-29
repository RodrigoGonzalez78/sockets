package server

import (
	"encoding/json"
	"net"
	"time"

	"github.com/RodrigoGonzalez78/sockets_messages/models"
)

func sendMessageToClient(transmition models.Transmition, sender net.Conn) {

	client, ok := clients[transmition.Messaje.IP]
	transmition.Time = time.Now().Format("15:04")

	if !ok {

		transmition.Operation = "error_message"
		transmition.Messaje.Contend = "Error el user " + transmition.Messaje.IP + " no esta disponible."
		json.NewEncoder(sender).Encode(transmition)
		return
	}

	transmition.Operation = "destination_receive_message"

	json.NewEncoder(sender).Encode(transmition)

	transmition.Operation = "new_message_received"
	transmition.Messaje.IP = sender.RemoteAddr().String()

	json.NewEncoder(client.Connection).Encode(transmition)
}
