package file_manager

import (
	"encoding/csv"
	"fmt"
	"os"
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
