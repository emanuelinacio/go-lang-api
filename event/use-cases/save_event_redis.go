package usecases

import (
	e "events/event/entity"
	r "events/event/repository"
)

func SaveEventRedis(event e.Event) (bool, error) {

	repository := new(r.RedisRepository)
	return SaveEvent(event, repository)
}
