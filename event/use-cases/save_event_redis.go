package usecases

import (
	e "events/event/entity"
	r "events/event/external"
)

func SaveEventRedis(event e.Event) (bool, error) {

	repository := new(r.RedisRepository)
	return SaveEvent(event, repository)
}
