package handlers

import (
	"context"
	"encoding/json"
	"escort-book-escort-consumer/models"
	"escort-book-escort-consumer/repositories"
	"log"

	"github.com/confluentinc/confluent-kafka-go/kafka"
)

type EscortHandler struct {
	ProfileRepository               repositories.IProfileRepository
	ProfileStatusRepository         repositories.IProfileStatusRepository
	ProfileStatusCategoryRepository repositories.IProfileStatusCategoryRepository
	NationalityRepository           repositories.INationalityRepository
}

func (h *EscortHandler) ProcessMessage(ctx context.Context, message *kafka.Message) {
	var profile models.Profile
	value := message.Value

	json.Unmarshal(value, &profile)
	profile.SetDefaultValues()

	nationality, err := h.NationalityRepository.GetOneByName(ctx, "empty")

	if err != nil {
		log.Println("ERROR GETTING THE NATIONALITY: ", err.Error())
		return
	}

	profile.NationalityId = nationality.Id
	err = h.ProfileRepository.Create(ctx, &profile)

	if err != nil {
		log.Println("ERROR INSERTING THE PROFILE: ", err.Error())
		return
	}

	status, err := h.ProfileStatusCategoryRepository.GetOneByName(ctx, "In Review")

	if err != nil {
		log.Println("ERROR GETTING THE PROFILE STATUS: ", err.Error())
		return
	}

	profileStatus := models.ProfileStatus{
		EscortId:                profile.EscortId,
		ProfileStatusCategoryId: status.Id,
	}
	profileStatus.SetDefaultValues()
	err = h.ProfileStatusRepository.Create(ctx, &profileStatus)

	if err != nil {
		log.Println("ERROR INSERTING THE PROFILE STATUS: ", err.Error())
	}
}
