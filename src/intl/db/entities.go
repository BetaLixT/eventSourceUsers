package db

import (
	"encoding/json"
	"time"

)

type Code struct {
	Code             string     `db:"code"`
	Title            string     `db:"title"`
	Description      string     `db:"description"`
	RedeemType       string     `db:"redeemtype"`
	Link             string     `db:"link"`
	IsRedeemed       bool       `db:"isredeemed"`
	DateTimeRedeemed *time.Time `db:"datetimeredeemed"`
	DateTimeCreated  *time.Time `db:"datetimecreated"`
	DateTimeUpdated  *time.Time `db:"datetimeupdated"`
}

type User struct {
	Id              string     `db:"id"`
	Provider        string     `db:"provider"`
	Roles           string     `db:"roles"`
	DateTimeCreated *time.Time `db:"datetimecreated"`
	DateTimeUpdated *time.Time `db:"datetimeupdated"`
}

func (usr *User) GetRoles() (roles []string) {
	json.Unmarshal([]byte(usr.Roles), &roles)
	return
}

type ExistsEntity struct {
	Exists bool `db:"exists"`
}

type CountEntity struct {
	Count int `db:"count"`
}