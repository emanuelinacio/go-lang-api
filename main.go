package main

import (
	e "events/event/entity"
	"fmt"
)

func main() {

	//localEvent := e.Event{"EventType", "Source", "Topic", `{"referer":"/golang"}`}
	eventStruct := new(e.Event)
	localEvent, err := eventStruct.Init("view", "Source", "Topic", `{"referer":/golang"}`)

	if err == nil {
		fmt.Println(localEvent.ToJson())
	}

	fmt.Println(err)
}
