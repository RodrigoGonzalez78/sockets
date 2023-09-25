package models

import (
	"net"
	"time"
)

type Log struct {
	Direccion net.Addr
	Fecha     time.Time
	Operacion string
}

func (l Log) ConvertirAString() []string {
	return []string{l.Direccion.String(), l.Operacion, l.Fecha.String()}
}
