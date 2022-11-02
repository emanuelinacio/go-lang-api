package usecases

import (
	e "events/event/entity"
)

func CreateEvent(Type string, Source string, Topic string, Extras string) (e.Event, error) {

	createEvent, errCreateEvent := e.NewEvent(Type, Source, Topic, Extras)

	if errCreateEvent != nil {
		return e.Event{}, errCreateEvent
	}

	return createEvent, nil
}
