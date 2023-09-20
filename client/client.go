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
			var mensaje models.Mensaje

			err := json.NewDecoder(conn).Decode(&mensaje)

			if err != nil {
				fmt.Println("Error al leer el mensaje del servidor:", err)
				return
			}

			fmt.Printf("\n## %s : %s  %v:%v##\n", mensaje.NombreCliente, mensaje.Mensaje, mensaje.FechaHora.Hour(), mensaje.FechaHora.Minute())
		}
	}()

	for {
		// Leer el mensaje del usuario
		textoMensaje, _ := reader.ReadString('\n')

		// Crear un mensaje
		mensaje := models.Mensaje{
			NombreCliente: nombreCliente,
			Mensaje:       strings.TrimSpace(textoMensaje),
			FechaHora:     time.Now(),
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

		writer.WriteString("\n") // Agregar una nueva línea para indicar el fin del mensaje
		writer.Flush()
	}
}
