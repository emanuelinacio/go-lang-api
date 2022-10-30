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

func (e Event) Init(Type string, Source string, Topic string, Extras string) (Event, error) {

	objectValueType := new(oV.ObType)
	validTypeORError, errType := objectValueType.Init(Type)

	if errType != nil {
		return e, errType
	}

	e.EventType = validTypeORError

	objectValueExtras := new(oV.ObExtras)
	validExtrasOrError, errExtras := objectValueExtras.Init(Extras)

	if errExtras != nil {
		return e, errExtras
	}

	e.Extras = validExtrasOrError

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
