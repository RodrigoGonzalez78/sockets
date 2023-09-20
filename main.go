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

	file_manager.CrearArchivoCSV("/logs.cvs")
	ip := "127.0.0.1"
	puerto := "8080"
	direccion := ip + ":" + puerto

	fmt.Println("Bienvenido al programa de servidor/cliente!")
	fmt.Println("1. Iniciar servidor")
	fmt.Println("2. Iniciar cliente")

	reader := bufio.NewReader(os.Stdin)
	fmt.Print("Elije una opción: ")

	opcionStr, err := reader.ReadString('\n')
	if err != nil {
		fmt.Println("Error al leer la entrada:", err)
		return
	}

	opcion := strings.TrimSpace(opcionStr)

	switch opcion {
	case "1":
		fmt.Println("Iniciando servidor en", direccion)
		// Llamar a la función para iniciar el servidor
		server.StartServer(direccion)
	case "2":
		fmt.Print("Ingresa tu nombre: ")
		nombreCliente, err := reader.ReadString('\n')
		if err != nil {
			fmt.Println("Error al leer la entrada:", err)
			return
		}
		nombreCliente = strings.TrimSpace(nombreCliente)
		fmt.Println("Iniciando cliente con nombre", nombreCliente, "en", direccion)
		// Llamar a la función para iniciar el cliente
		client.StartClient(direccion, nombreCliente)
	default:
		fmt.Println("Opción no válida. Saliendo.")
	}

}
