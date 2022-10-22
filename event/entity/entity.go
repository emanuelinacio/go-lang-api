package entity

import (
	"errors"
    "fmt"
	"encoding/json"
)

type Event struct{
	EventType, Source, Topic, Extras string
}

func (e Event) ToJson() (string, error){
	
	JsonObject, err := json.Marshal(e)

	if err != nil {
		fmt.Println(err)
		return "", errors.New("json convert failed")
	}

	return string(JsonObject), nil
}

func Sum(x int, y int) int{
	return x + y
}
