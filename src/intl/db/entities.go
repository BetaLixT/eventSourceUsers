package db

import (
)

type ExistsEntity struct {
	Exists bool `db:"exists"`
}

type CountEntity struct {
	Count int `db:"count"`
}
