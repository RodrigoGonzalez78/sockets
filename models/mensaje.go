package models

import "time"

type Mensaje struct {
	NombreCliente string
	Mensaje       string
	FechaHora     time.Time
}

func (m Mensaje) ConvertirAString() []string {
	return []string{m.Mensaje, m.NombreCliente, m.FechaHora.String()}
}
