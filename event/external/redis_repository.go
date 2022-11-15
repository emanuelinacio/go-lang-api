package external

import (
	"context"
	"errors"
	e "events/event/entity"

	redis "github.com/go-redis/redis/v8"
)

type RedisRepository struct {
	Event e.Event
}

func (rR RedisRepository) SaveEvent(event e.Event) (bool, error) {
	rR.Event = event
	return rR.SaveEventRedis()
}

func (rR RedisRepository) SaveEventRedis() (bool, error) {

	ctx := context.Background()

	rdb := redis.NewClient(&redis.Options{
		Addr:     "redis:6379",
		Password: "",
		DB:       0,
	})

	json, errJson := rR.Event.ToJson()

	if errJson != nil {
		return false, errors.New(errJson.Error())
	}

	err := rdb.Set(ctx, "event", json, 0).Err()

	if err != nil {
		return false, errors.New(err.Error())
	}

	return true, nil
}
