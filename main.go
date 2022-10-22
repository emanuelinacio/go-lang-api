package main

import (
	"fmt"
	e "events/event/entity"
)

func main() {

	localEvent := e.Event{ "EventType", "Source", "Topic", "Extras" }

	fmt.Println( localEvent.ToJson() )
}