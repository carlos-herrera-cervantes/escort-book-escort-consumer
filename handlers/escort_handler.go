package handlers

import (
	"context"
	"encoding/json"
	"escort-book-escort-consumer/models"
	"escort-book-escort-consumer/repositories"
	"escort-book-escort-consumer/types"
	"log"
	"os"

	"github.com/confluentinc/confluent-kafka-go/kafka"
)

type EscortHandler struct {
	ProfileRepository               repositories.IProfileRepository
	ProfileStatusRepository         repositories.IProfileStatusRepository
	ProfileStatusCategoryRepository repositories.IProfileStatusCategoryRepository
	NationalityRepository           repositories.INationalityRepository
}

func (h *EscortHandler) ProcessMessage(ctx context.Context, message *kafka.Message) {
	topic := message.TopicPartition.Topic

	if *(topic) == os.Getenv("KAFKA_ACTIVE_ACCOUNT_TOPIC") {
		h.activeProfile(ctx, message)
		return
	}

	h.createProfile(ctx, message)
}

func (h *EscortHandler) activeProfile(ctx context.Context, message *kafka.Message) {
	var activeAccountEvent types.ActiveAccountEvent
	value := message.Value

	_ = json.Unmarshal(value, &activeAccountEvent)

	if _, err := h.ProfileStatusRepository.GetByProfileId(
		ctx,
		activeAccountEvent.UserId,
	); err != nil {
		return
	}

	category, err := h.ProfileStatusCategoryRepository.GetOneByName(ctx, "Active")

	if err != nil {
		log.Println("ERROR GETTING THE CATEGORY: ", err.Error())
		return
	}

	if err = h.ProfileStatusRepository.UpdateByProfileId(
		ctx,
		activeAccountEvent.UserId,
		category.Id,
	); err != nil {
		log.Println("ERROR UPDATING THE PROFILE STATUS: ", err.Error())
	}
}

func (h *EscortHandler) createProfile(ctx context.Context, message *kafka.Message) {
	var profile models.Profile
	value := message.Value

	_ = json.Unmarshal(value, &profile)
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
