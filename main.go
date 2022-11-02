package main

import (
	uc "events/event/use-cases"
	"fmt"
)

func main() {
	localEvent, err := uc.CreateEvent("view", "Source", "Topic", `{"referer":"/golang"}`)

	if err != nil {
		fmt.Println(err)
	}

	saveEvent, errSaveEvent := uc.SaveEventRedis(localEvent)

	if errSaveEvent != nil {
		fmt.Println(errSaveEvent.Error())
	}

	fmt.Println(localEvent.ToJson())
	fmt.Println(saveEvent)
}
