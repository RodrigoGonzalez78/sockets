package client

import (
	"bufio"
	"fmt"
	"net"
	"os"
)

func StartClient() {
	conn, err := net.Dial("tcp", "localhost:8080")
	if err != nil {
		fmt.Println("Error al conectar al servidor:", err)
		return
	}
	defer conn.Close()

	fmt.Println("Conexión establecida. Puedes empezar a enviar mensajes.")

	// Crear un lector para leer mensajes del usuario
	reader := bufio.NewReader(os.Stdin)

	// Crear un escritor para enviar mensajes al servidor
	writer := bufio.NewWriter(conn)

	for {
		// Leer el mensaje del usuario
		fmt.Print("Tú: ")
		message, _ := reader.ReadString('\n')

		// Enviar el mensaje al servidor
		_, err := writer.WriteString(message)
		if err != nil {
			fmt.Println("Error al enviar mensaje al servidor:", err)
			return
		}
		writer.Flush()

		// Leer la respuesta del servidor
		response, err := bufio.NewReader(conn).ReadString('\n')
		if err != nil {
			fmt.Println("Error al leer la respuesta del servidor:", err)
			return
		}

		fmt.Println("Servidor:", response)
	}
}
