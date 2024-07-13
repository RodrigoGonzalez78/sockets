package client

import (
	"bufio"
	"encoding/json"
	"fmt"
	"net"
	"os"
	"strings"
	"time"

	"github.com/RodrigoGonzalez78/sockets_messages/models"
)

func StartClient(direccion, nombreCliente string) {

	//Inicio una nueva conexion al servidor
	conn, err := net.Dial("tcp", direccion)

	if err != nil {
		fmt.Println("Error al conectar al servidor con la direccion:", direccion, err)
		return
	}

	defer conn.Close()

	fmt.Println("Conexión establecida. Puedes empezar a enviar mensajes.")

	// Crear un lector para leer mensajes del usuario
	reader := bufio.NewReader(os.Stdin)

	// Crear un escritor para enviar mensajes al servidor
	writer := bufio.NewWriter(conn)

	// Crear una goroutine para recibir mensajes y mostrarlos en la consola
	go func() {
		for {
			var mensaje models.Message

			err := json.NewDecoder(conn).Decode(&mensaje)

			if err != nil {

				//Un error de entrada/salida significa que se cerro la conexion
				if err.Error() == "EOF" {
					os.Exit(0)
				}
				fmt.Println("Error al leer el mensaje del servidor:", err)
				return
			}

			fmt.Printf("\n## %s : %s  %v##\n", mensaje.ClientName, mensaje.Message, mensaje.Time)
		}
	}()

	for {
		// Leer el mensaje del usuario
		textoMensaje, _ := reader.ReadString('\n')

		// Crear un mensaje
		mensaje := models.Message{
			ClientName: nombreCliente,
			Message:    strings.TrimSpace(textoMensaje),
			Time:       time.Now().Format("15:04"),
		}

		// Convertir el mensaje a JSON
		jsonData, err := json.Marshal(mensaje)
		if err != nil {
			fmt.Println("Error al convertir mensaje a JSON:", err)
			return
		}

		// Enviar el mensaje al servidor
		_, err = writer.Write(jsonData)

		if err != nil {
			fmt.Println("Error al enviar mensaje al servidor:", err)
			return
		}

		writer.WriteString("\n")
		writer.Flush()
	}
}
