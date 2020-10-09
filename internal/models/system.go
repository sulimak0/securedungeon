package models

type System struct {
	state bool
}

func (s *System) State() bool {
	return s.state
}
