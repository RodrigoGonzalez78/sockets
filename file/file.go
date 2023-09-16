package file

import (
	"encoding/csv"
	"os"
)

func File() {
	// Abrir un archivo para escritura
	file, err := os.Create("datos.csv")
	if err != nil {
		panic(err)
	}
	defer file.Close()

	// Crear un escritor CSV
	writer := csv.NewWriter(file)
	defer writer.Flush()

	// Datos que queremos escribir en el archivo CSV
	data := [][]string{
		{"Nombre", "Edad", "Ciudad"},
		{"Alice", "30", "Nueva York"},
		{"Bob", "35", "Los Ángeles"},
		{"Charlie", "25", "Chicago"},
	}

	// Escribir los datos en el archivo CSV
	for _, record := range data {
		if err := writer.Write(record); err != nil {
			panic(err)
		}
	}

	// Flush para asegurarse de que todos los datos estén escritos en el archivo
	writer.Flush()

	if err := writer.Error(); err != nil {
		panic(err)
	}

	println("Datos escritos en datos.csv")
}
