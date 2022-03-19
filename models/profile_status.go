package models

import (
	"time"

	"github.com/google/uuid"
)

type ProfileStatus struct {
	Id                      string
	EscortId                string
	ProfileStatusCategoryId string
	CreatedAt               time.Time
	UpdatedAt               time.Time
}

func (p *ProfileStatus) SetDefaultValues() {
	p.Id = uuid.NewString()
}
