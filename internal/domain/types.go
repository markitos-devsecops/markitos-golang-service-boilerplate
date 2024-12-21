package domain

type BoilerId struct {
	value string
}

func NewBoilerId(value string) (*BoilerId, error) {
	if value == "" {
		return nil, NewBoilerInvalidIdError(value)
	}
	if !IsUUIDv4(value) {
		return nil, NewBoilerInvalidIdFormatError(value)
	}

	return &BoilerId{value: value}, nil
}

func (id *BoilerId) Value() string {
	return id.value
}

type BoilerMessage struct {
	value string
}

func NewBoilerMessage(value string) (*BoilerMessage, error) {
	if value == "" {
		return nil, NewBoilerInvalidMessageError(value)
	}
	return &BoilerMessage{value: value}, nil
}

func (msg *BoilerMessage) Value() string {
	return msg.value
}
