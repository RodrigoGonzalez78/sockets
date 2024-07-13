package models

type Mensaje struct {
	ClientName string
	Message    string
	Time       string
}

func (m Mensaje) ToString() []string {
	return []string{m.ClientName, m.Message, m.Time}
}
