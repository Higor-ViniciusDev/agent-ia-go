package events

import "time"

type WorkCreated struct {
	Name   string
	Values any
}

func NewWorkCreated() *WorkCreated {
	return &WorkCreated{
		Name: "WorkCreated",
	}
}

func (e *WorkCreated) GetName() string {
	return e.Name
}

func (e *WorkCreated) GetPayload() any {
	return e.Values
}

func (e *WorkCreated) SetPayload(values any) {
	e.Values = values
}

func (e *WorkCreated) GetDateTime() time.Time {
	return time.Now()
}
