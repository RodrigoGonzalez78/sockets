package file_manager

import (
	"encoding/csv"
	"fmt"
	"os"
	"time"

	"github.com/RodrigoGonzalez78/sockets_messages/models"
)

func CrearArchivoCSV(ruta string) error {
	// Verificar si el archivo ya existe
	_, err := os.Stat(ruta)

	if err == nil {
		// El archivo ya existe, no es necesario crearlo
		fmt.Println("El archivo CSV ya existe en la ruta especificada.")
		return nil
	}

	// Si os.IsNotExist devuelve true, entonces el archivo no existe y lo creamos
	if os.IsNotExist(err) {
		file, err := os.Create(ruta)
		if err != nil {
			return err
		}
		defer file.Close()
		fmt.Println("Archivo CSV creado en la ruta especificada.")
		return nil
	}

	// Otro error, retornamos el error
	return err
}

func EscribirDatosEnCSV(ruta string, informacion []string) error {

	file, err := os.OpenFile(ruta, os.O_WRONLY|os.O_APPEND, os.ModeAppend)
	if err != nil {
		return err
	}
	defer file.Close()

	writer := csv.NewWriter(file)
	defer writer.Flush()

	if err := writer.Write(informacion); err != nil {
		return err
	}

	return nil
}

func LeerDatosDeCSV(ruta string) ([]models.Mensaje, error) {
	file, err := os.Open(ruta)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	reader := csv.NewReader(file)
	records, err := reader.ReadAll()
	if err != nil {
		return nil, err
	}

	var mensajes []models.Mensaje

	for _, record := range records {
		if len(record) >= 3 {

			// Parsear la fecha y hora
			fechaHora, err := time.Parse(time.RFC3339, record[2])
			if err != nil {
				return nil, err
			}

			// Crear un nuevo mensaje
			mensaje := models.Mensaje{
				Mensaje:       record[0],
				NombreCliente: record[1],
				FechaHora:     fechaHora,
			}

			mensajes = append(mensajes, mensaje)
		}
	}

	return mensajes, nil
}
