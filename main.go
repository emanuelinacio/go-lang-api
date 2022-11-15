package main

import (
	c "events/event/presentation/controllers/web"
	"fmt"
	"net/http"
)

func hello(res http.ResponseWriter, req *http.Request) {

	if req.Method == "GET" {
		res.WriteHeader(401)
		fmt.Fprintf(res, "Method Request Error")
	}
}

func main() {

	controller := c.NewSendEventController()
	controller.SetRouter()
	fmt.Println("Stating server 8080")
	http.ListenAndServe(":8080", nil)

	/*localEvent, err := uc.CreateEvent("view", "Source", "Topic", `{"referer":"/golang"}`)

	if err != nil {
		fmt.Println(err)
	}

	saveEvent, errSaveEvent := uc.SaveEventRedis(localEvent)

	if errSaveEvent != nil {
		fmt.Println(errSaveEvent.Error())
	}

	fmt.Println(localEvent.ToJson())
	fmt.Println(saveEvent)*/
}
