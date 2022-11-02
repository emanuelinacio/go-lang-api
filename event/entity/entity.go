package entity

import (
	"encoding/json"
	"errors"
	oV "events/event/entity/object_values"
	"fmt"
)

type Event struct {
	EventType     oV.ObType
	Extras        oV.ObExtras
	Source, Topic string
}

func NewEvent(Type string, Source string, Topic string, Extras string) (Event, error) {

	e := new(Event)
	eventValidated, errValidation := e.ValidInput(Type, Source, Topic, Extras)

	if errValidation != nil {
		fmt.Println(errValidation)
		return Event{}, errValidation
	}

	return eventValidated, nil
}

func (e Event) ValidInput(Type string, Source string, Topic string, Extras string) (Event, error) {

	validType, errType := oV.NewObType(Type)

	if errType != nil {
		return e, errType
	}

	e.EventType = validType
	validExtras, errExtras := oV.NewObExtras(Extras)

	if errExtras != nil {
		return e, errExtras
	}

	e.Extras = validExtras
	e.Source = Source
	e.Topic = Topic

	return e, nil
}

func (e Event) ToJson() (string, error) {

	JsonObject, err := json.Marshal(e)

	if err != nil {
		fmt.Println(err)
		return "", errors.New("json convert failed")
	}

	return string(JsonObject), nil
}
