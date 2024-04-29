package server

import (
	"encoding/json"
	"net"
	"time"

	"github.com/RodrigoGonzalez78/sockets_messages/models"
)

func sendConnectedIPs(conn net.Conn) {
	var connectedIPs []string
	for addr := range clients {
		connectedIPs = append(connectedIPs, addr)
	}

	transmition := models.Transmition{
		Operation: "get_clients_list",
		Time:      time.Now().Format("15:04"),
	}
	json.NewEncoder(conn).Encode(transmition)
}
