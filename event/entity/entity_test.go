package entity

import (
	"fmt"
	"testing"
)

func TestIsValidEventType(t *testing.T) {

	testExtrasEvent := string(`{"referer":"/golang"}`)

	eventStruct := new(Event)
	eventTest, err := eventStruct.Init("vieww", "Source", "Topic", testExtrasEvent)

	if err == nil {
		t.Error("Expect to be false", err)
	}

	fmt.Println(eventTest.ToJson())
}

func TestIsValidExtras(t *testing.T) {

	testExtrasEvent := string(`{"referer":/golang"}`)

	eventStruct := new(Event)
	eventTest, err := eventStruct.Init("view", "Source", "Topic", testExtrasEvent)

	if err == nil {
		t.Error("Expect to be false", err)
	}

	fmt.Println(eventTest.ToJson())
}
