package main

import (
	e "events/event/entity"
	"fmt"
)

func main() {

	localEvent, err := e.NewEvent("view", "Source", "Topic", `{"referer":"/golang"}`)

	if err == nil {
		fmt.Println(localEvent.ToJson())
	}

	fmt.Println(err)
}
