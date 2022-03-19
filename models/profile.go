package models

import (
	"time"

	"github.com/google/uuid"
)

type Profile struct {
	Id            string
	EscortId      string `json:"_id"`
	FirstName     string
	LastName      string
	Email         string `json:"email"`
	PhoneNumber   string
	Gender        string
	NationalityId string
	Birthdate     string
	CreatedAt     time.Time
	UpdatedAt     time.Time
}

func (p *Profile) SetDefaultValues() {
	p.Id = uuid.NewString()
	p.Gender = "NotSpecified"
}
