package models

type Mensaje struct {
	NombreCliente string
	Mensaje       string
	FechaHora     string
}

func (m Mensaje) ConvertirAString() []string {
	return []string{m.NombreCliente, m.Mensaje, m.FechaHora}
}
