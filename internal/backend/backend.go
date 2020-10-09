package backend

import "github.com/sulimak0/securedungeon/internal/models"

type SecuritySystemUnit struct {
	state    bool
	url      string
	messages chan models.Alarm
}

func NewSecuritySystem(url string, messages chan models.Alarm) *SecuritySystemUnit {
	return &SecuritySystemUnit{state: false, url: url, messages: messages}
}

func (s *SecuritySystemUnit) TurnOn() {
	s.state = true
}

func (s *SecuritySystemUnit) TurnOff() {
	s.state = false
}

func (s *SecuritySystemUnit) Url() string {
	return s.url
}

func (s *SecuritySystemUnit) State() bool {
	return s.state
}

func (s *SecuritySystemUnit) GetMessage() models.Alarm {
	select {
	case m := <-s.messages:
		return m
	}
	return models.Alarm{}
}
