package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"

	"github.com/RodrigoGonzalez78/sockets_messages/server"
)

func main() {
	fmt.Println("Bienvenido al programa de mensajería por sockets.")

	reader := bufio.NewReader(os.Stdin)

	direccion, err := cargarDireccion(reader)

	if err != nil {
		return
	}

	fmt.Println("Iniciando servidor en ", direccion)
	server.StartServer(direccion)

}

func cargarDireccion(reader *bufio.Reader) (string, error) {
	fmt.Print("Ingresa la dirección IP: ")
	ip, err := reader.ReadString('\n')

	if err != nil {
		fmt.Println("Error al leer la entrada:", err)
		return "", err
	}

	ip = strings.TrimSpace(ip)

	fmt.Print("Ingresa el puerto: ")
	puerto, err := reader.ReadString('\n')

	if err != nil {
		fmt.Println("Error al leer la entrada:", err)
		return "", err
	}

	puerto = strings.TrimSpace(puerto)
	return ip + ":" + puerto, nil
}
