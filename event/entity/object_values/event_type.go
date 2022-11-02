package event_type

import (
	"errors"
)

type ObType struct {
	Type string
}

func NewObType(Type string) (ObType, error) {
	o := new(ObType)
	o.Type = Type

	if o.IsValidEventType() {
		return *o, nil
	}

	return ObType{}, errors.New("Invalid type")
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
