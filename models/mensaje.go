package models

type Message struct {
	ClientName string
	Message    string
	Time       string
}

func (m Message) ToString() []string {
	return []string{m.ClientName, m.Message, m.Time}
}
