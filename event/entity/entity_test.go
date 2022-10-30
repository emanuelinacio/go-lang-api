package entity

import (
	"fmt"
	"testing"
)

func TestIsValidEventType(t *testing.T) {

	testExtrasEvent := string(`{"referer":"/golang"}`)
	eventTest, err := NewEvent("view", "Source", "Topic", testExtrasEvent)

	if err != nil {
		t.Error("Expect to be true", err)
	}

	fmt.Println(eventTest.ToJson())
}

func TestIsValidExtras(t *testing.T) {

	testExtrasEvent := string(`{"referer":"/golang"}`)
	eventTest, err := NewEvent("view", "Source", "Topic", testExtrasEvent)

	if err != nil {
		t.Error("Expect to be true", err)
	}

	fmt.Println(eventTest.ToJson())
}
