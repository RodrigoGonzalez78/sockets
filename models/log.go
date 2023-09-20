package models

import "time"

type Log struct {
	Direccion string
	Fecha     time.Time
	Operacion string
}

func (l Log) ConvertirAString() []string {
	return []string{l.Direccion, l.Fecha.String(), l.Operacion}
}
