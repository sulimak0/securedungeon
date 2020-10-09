package models

type Alarm struct {
	Counter int
	Message string
}

func New(message string) *Alarm {
	return &Alarm{Message: message}
}
