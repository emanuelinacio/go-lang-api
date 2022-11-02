package usecases

import (
	e "events/event/entity"
	r "events/event/repository"
)

func SaveEvent(event e.Event, repository r.Repository) (bool, error) {
	return repository.SaveEvent(event)
}
