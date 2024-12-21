package domain

import (
	"fmt"
)

// --------------------------------------------------------------
const BOILER_INVALID_ID_PREFIX = "invalid id"

type BoilerInvalidIdError struct {
	id string
}

func NewBoilerInvalidIdError(id string) error {
	return &BoilerInvalidIdError{id: id}
}

func (e *BoilerInvalidIdError) Error() string {
	return fmt.Sprintf("%s: %s", BOILER_INVALID_ID_PREFIX, e.id)
}

// --------------------------------------------------------------
const BOILER_INVALID_MESSAGE_PREFIX = "invalid message"

type BoilerInvalidMessageError struct {
	message string
}

func NewBoilerInvalidMessageError(message string) error {
	return &BoilerInvalidMessageError{message: message}
}

func (e *BoilerInvalidMessageError) Error() string {
	return fmt.Sprintf("%s: %s", BOILER_INVALID_MESSAGE_PREFIX, e.message)
}

// --------------------------------------------------------------
const BOILER_INVALID_ID_FORMAT_PREFIX = "invalid id format, must be an UUIDv4"

type BoilerInvalidIdFormatError struct {
	id string
}

func NewBoilerInvalidIdFormatError(id string) error {
	return &BoilerInvalidIdFormatError{id: id}
}

func (e *BoilerInvalidIdFormatError) Error() string {
	return fmt.Sprintf("%s: %s", BOILER_INVALID_ID_FORMAT_PREFIX, e.id)
}

//--------------------------------------------------------------
