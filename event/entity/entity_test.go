package entity

import (
	"testing"
)

func TestIsValidEventType(t *testing.T) {

	testExtrasEvent := string(`{"referer":"/golang"}`)
	testEvent := Event{"EventType", "Source", "Topic", testExtrasEvent}

	if testEvent.IsValidEventType() == true {
		t.Error("Excpet to be false ", testEvent.EventType)
	}
}

func TestIsValidExtras(t *testing.T) {

	testExtrasEvent := string(`{"Referer":golang"}`)
	testEvent := Event{"EventType", "Source", "Topic", testExtrasEvent}

	if testEvent.IsValidExtras() == true {
		t.Error("Excpet to be false ", testEvent.Extras)
	}

}
