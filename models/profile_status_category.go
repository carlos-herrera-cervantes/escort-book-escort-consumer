package models

import "time"

type ProfileStatusCategory struct {
	Id        string
	Name      string
	Active    bool
	CreatedAt time.Time
	UpdatedAt time.Time
}
