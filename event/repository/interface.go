package repository

import e "events/event/entity"

type Repository interface {
	SaveEvent(e.Event) (bool, error)
}
