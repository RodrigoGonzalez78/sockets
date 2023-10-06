package server

import (
	"bufio"
	"encoding/json"
	"fmt"
	"net"
	"time"

	"github.com/RodrigoGonzalez78/sockets_messages/file_manager"
	"github.com/RodrigoGonzalez78/sockets_messages/models"
)

var clients []models.Client

var LogsFile string = "logs.csv"

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
			saveLog("Error al aceptar la conexión: "+err.Error(), conn)
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

	saveLog("Nueva Conexion", conn)

	reader := bufio.NewReader(conn)

	for {
		var mensaje models.Mensaje

		err := json.NewDecoder(reader).Decode(&mensaje)

		if err != nil {
			fmt.Println("Error al leer el mensaje del cliente:", err)
			saveLog("Error al leer el mensaje del cliente:"+err.Error(), conn)
			removeClient(conn)
			return
		}

		fmt.Printf("\n## %s : %s %v ##\n", mensaje.NombreCliente, mensaje.Mensaje, mensaje.FechaHora)

		if mensaje.Mensaje == "/listar" {
			sendConnectedClientsList(conn)
		} else if mensaje.Mensaje == "/quitar" {
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

	listaClientes := models.Mensaje{
		NombreCliente: "Servidor",
		Mensaje:       clientesString,
		FechaHora:     time.Now().Format("15:04"),
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

	saveLog("Cliente desconectado", conn)

}

func saveLog(log string, conn net.Conn) {
	newLog := models.Log{
		Direccion: conn.RemoteAddr(),
		Fecha:     time.Now(),
		Operacion: log,
	}

	file_manager.EscribirDatosEnCSV(LogsFile, newLog.ConvertirAString())
}

func sendDisconnectMessage(conn net.Conn) {
	cerrarConexion := models.Mensaje{
		NombreCliente: "Servidor",
		Mensaje:       "Tu sesión se ha cerrado.",
		FechaHora:     time.Now().Format("15:04"),
	}
	json.NewEncoder(conn).Encode(cerrarConexion)
	removeClient(conn)
	conn.Close()
}

func broadcastMessage(mensaje models.Mensaje, direccion net.Addr) {

	for _, client := range clients {
		if client.Connection.RemoteAddr() != direccion {
			json.NewEncoder(client.Connection).Encode(mensaje)
		}

	}
}
