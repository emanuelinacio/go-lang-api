package entity

import (
	"encoding/json"
	"errors"
	"fmt"
	//"strings"
)

type Event struct {
	EventType, Source, Topic, Extras string
}

func (e Event) ToJson() (string, error) {

	JsonObject, err := json.Marshal(e)

	if err != nil {
		fmt.Println(err)
		return "", errors.New("json convert failed")
	}

	return string(JsonObject), nil
}

func (e Event) IsValidEventType() bool {

	switch e.EventType {
	case "view":
		return true
	case "impression":
		return true
	case "conversion":
		return true
	}

	return false
}

func (e Event) IsValidExtras() bool {

	var dat map[string]interface{}
	err := json.Unmarshal([]byte(e.Extras), &dat)

	return err == nil
}
