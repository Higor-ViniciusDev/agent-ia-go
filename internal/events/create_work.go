package events

import "time"

type WorkCreated struct {
	Nome   string
	Values any
}

func NewWorkCreated() *WorkCreated {
	return &WorkCreated{
		Nome: "WorkCreated",
	}
}

func (e *WorkCreated) GetNome() string {
	return e.Nome
}

func (e *WorkCreated) GetValues() any {
	return e.Values
}

func (e *WorkCreated) SetValues(values any) {
	e.Values = values
}

func (e *WorkCreated) GetDateTime() time.Time {
	return time.Now()
}
