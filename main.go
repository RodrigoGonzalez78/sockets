package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"

	"github.com/RodrigoGonzalez78/sockets_messages/client"
	"github.com/RodrigoGonzalez78/sockets_messages/file_manager"
	"github.com/RodrigoGonzalez78/sockets_messages/server"
)

func main() {
	fmt.Println("Bienvenido al programa de mensajería por sockets.")
	fmt.Println("1. Iniciar servidor")
	fmt.Println("2. Iniciar cliente")

	reader := bufio.NewReader(os.Stdin)
	fmt.Print("Elije una opción: ")

	opcionStr, err := reader.ReadString('\n')
	if err != nil {
		fmt.Println("Error al leer la entrada de la opcion:", err)
		return
	}

	opcion := strings.TrimSpace(opcionStr)

	direccion, err := cargarDireccion(reader)

	if err != nil {
		return
	}

	switch opcion {
	case "1":
		opcionServidor(direccion)
	case "2":
		opcionCliente(direccion, reader)
	default:
		fmt.Println("Opción no válida. Saliendo.")
	}
}

func opcionServidor(direccion string) {
	fmt.Println("Iniciando servidor en", direccion)
	file_manager.CrearArchivoCSV(server.LogsFile)

	server.StartServer(direccion)
}

func opcionCliente(direccion string, reader *bufio.Reader) {
	fmt.Print("Ingresa tu nombre: ")
	nombreCliente, err := reader.ReadString('\n')

	if err != nil {
		fmt.Println("Error al leer la entrada:", err)
		return
	}

	nombreCliente = strings.TrimSpace(nombreCliente)
	fmt.Println("Iniciando cliente con nombre", nombreCliente, "en", direccion)
	client.StartClient(direccion, nombreCliente)
}

func cargarDireccion(reader *bufio.Reader) (string, error) {
	fmt.Print("Ingresa la dirección IP: ")
	ip, err := reader.ReadString('\n')
	if err != nil {
		fmt.Println("Error al leer la entrada:", err)
		return "", err
	}
	ip = strings.TrimSpace(ip)
	fmt.Print("Ingresa el puerto para el cliente: ")
	puerto, err := reader.ReadString('\n')
	if err != nil {
		fmt.Println("Error al leer la entrada:", err)
		return "", err
	}
	puerto = strings.TrimSpace(puerto)
	return ip + ":" + puerto, nil
}
