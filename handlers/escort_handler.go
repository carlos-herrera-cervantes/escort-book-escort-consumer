package handlers

import (
	"context"
	"encoding/json"
	"escort-book-escort-consumer/models"
	"escort-book-escort-consumer/repositories"

	"github.com/confluentinc/confluent-kafka-go/kafka"
)

type EscortHandler struct {
	ProfileRepository               *repositories.ProfileRepository
	ProfileStatusRepository         *repositories.ProfileStatusRepository
	ProfileStatusCategoryRepository *repositories.ProfileStatusCategoryRepository
}

func (h *EscortHandler) ProcessMessage(ctx context.Context, message *kafka.Message) {
	var profile models.Profile
	value := message.Value

	json.Unmarshal(value, &profile)
	profile.SetDefaultValues()

	h.ProfileRepository.Create(ctx, &profile)
	status, _ := h.ProfileStatusCategoryRepository.GetOneByName(ctx, "In Review")

	profileStatus := models.ProfileStatus{
		EscortId:                profile.EscortId,
		ProfileStatusCategoryId: status.Id,
	}
	profileStatus.SetDefaultValues()
	h.ProfileStatusRepository.Create(ctx, &profileStatus)
}
