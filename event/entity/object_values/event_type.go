package event_type

import (
	"errors"
)

type ObType struct {
	Type string
}

func (o ObType) Init(Type string) (ObType, error) {
	o.Type = Type

	if o.IsValidEventType() {
		return o, nil
	}

	return o, errors.New("Invalid type")
}

func (o ObType) IsValidEventType() bool {

	switch o.Type {
	case "view":
		return true
	case "impression":
		return true
	case "conversion":
		return true
	}

	return false
}
