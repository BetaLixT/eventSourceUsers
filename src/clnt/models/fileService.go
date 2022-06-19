package models

import "time"

type FileMeta struct {
	FileId          string    `json:"fileId"`
	DateTimeCreated time.Time `json:"DateTimeCreated"`
	Type            string    `json:"type"`
}
