package handlers

import (
	"context"
	"encoding/json"
	"fmt"

	"escort-book-escort-consumer/config"
	"escort-book-escort-consumer/models"
	"escort-book-escort-consumer/repositories"
	"escort-book-escort-consumer/types"

	"github.com/confluentinc/confluent-kafka-go/kafka"
	log "github.com/inconshreveable/log15"
)

type EscortHandler struct {
	ProfileRepository               repositories.IProfileRepository
	ProfileStatusRepository         repositories.IProfileStatusRepository
	ProfileStatusCategoryRepository repositories.IProfileStatusCategoryRepository
	NationalityRepository           repositories.INationalityRepository
}

var logger = log.New("handlers")

func (h *EscortHandler) HandleEvent(ctx context.Context, message *kafka.Message) {
	topic := message.TopicPartition.Topic

	if *(topic) == config.InitializeKafka().Topics.UserActiveAccount {
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
		logger.Error(fmt.Sprintf("ERROR GETTING PROFILE STATUS: %s", err.Error()))
		return
	}

	category, err := h.ProfileStatusCategoryRepository.GetOneByName(ctx, "Active")

	if err != nil {
		logger.Error(fmt.Sprintf("ERROR GETTING THE CATEGORY: %s", err.Error()))
		return
	}

	if err = h.ProfileStatusRepository.UpdateByProfileId(
		ctx,
		activeAccountEvent.UserId,
		category.Id,
	); err != nil {
		logger.Error(fmt.Sprintf("ERROR UPDATING THE PROFILE STATUS: %s", err.Error()))
	}
}

func (h *EscortHandler) createProfile(ctx context.Context, message *kafka.Message) {
	var profile models.Profile
	value := message.Value

	_ = json.Unmarshal(value, &profile)
	profile.SetDefaultValues()

	nationality, err := h.NationalityRepository.GetOneByName(ctx, "empty")

	if err != nil {
		logger.Error(fmt.Sprintf("ERROR GETTING THE NATIONALITY: %s", err.Error()))
		return
	}

	profile.NationalityId = nationality.Id

	if err = h.ProfileRepository.Create(ctx, &profile); err != nil {
		logger.Error(fmt.Sprintf("ERROR INSERTING THE PROFILE: %s", err.Error()))
		return
	}

	status, err := h.ProfileStatusCategoryRepository.GetOneByName(ctx, "In Review")

	if err != nil {
		logger.Error(fmt.Sprintf("ERROR GETTING THE PROFILE STATUS: %s", err.Error()))
		return
	}

	profileStatus := models.ProfileStatus{
		EscortId:                profile.EscortId,
		ProfileStatusCategoryId: status.Id,
	}
	profileStatus.SetDefaultValues()

	if err = h.ProfileStatusRepository.Create(ctx, &profileStatus); err != nil {
		logger.Error(fmt.Sprintf("ERROR INSERTING THE PROFILE STATUS: %s", err.Error()))
	}
}
