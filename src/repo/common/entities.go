package common

import (
	"time"

	"github.com/jmoiron/sqlx/types"
)

type ExistsEntity struct {
	Exists bool `db:"exists"`
}

type CountEntity struct {
	Count int `db:"count"`
}

type Event struct {
	Id            string
	Stream        string
	StreamId      string
	StreamVersion int
	Event         string
	Data          types.JSONText
	EventDateTime time.Time
}
