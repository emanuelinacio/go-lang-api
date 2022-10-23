package main

import (
	e "events/event/entity"
	"fmt"
)

func main() {

	localEvent := e.Event{"EventType", "Source", "Topic", `{"referer":"/golang"}`}

	fmt.Println(localEvent.ToJson())
}
