package domain

import (
	"time"
)

type Boiler struct {
	Id        string    `json:"id" binding:"required,uuid"`
	Message   string    `json:"message" binding:"required"`
	CreatedAt time.Time `json:"created_at" binding:"required,datetime" default:"now"`
	UpdatedAt time.Time `json:"updated_at" binding:"required,datetime" default:"now"`
}

func NewBoiler(id, message string) (*Boiler, error) {
	anId, anIdError := NewBoilerId(id)
	if anIdError != nil {
		return nil, anIdError
	}

	aMessage, aMessageError := NewBoilerMessage(message)
	if aMessageError != nil {
		return nil, aMessageError
	}

	return &Boiler{
		Id:        anId.value,
		Message:   aMessage.value,
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}, nil
}
