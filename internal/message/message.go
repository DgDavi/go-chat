package message

import "time"

type Message struct {
	Nome  string    `json:"nome"`
	Texto string    `json:"texto"`
	Hora  time.Time `json:"hora"`
}
