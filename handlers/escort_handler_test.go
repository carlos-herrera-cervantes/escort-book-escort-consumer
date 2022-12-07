package handlers

import (
	"context"
	"errors"
	"testing"

	"escort-book-escort-consumer/config"
	"escort-book-escort-consumer/models"
	mockRepositories "escort-book-escort-consumer/repositories/mocks"

	"github.com/confluentinc/confluent-kafka-go/kafka"
	"github.com/golang/mock/gomock"
)

func TestEscortHandlerHandleEvent(t *testing.T) {
	controller := gomock.NewController(t)

	mockProfileRepository := mockRepositories.NewMockIProfileRepository(controller)
	mockProfileStatusRepository := mockRepositories.NewMockIProfileStatusRepository(controller)
	mockProfileStatusCategoryRepository := mockRepositories.NewMockIProfileStatusCategoryRepository(controller)
	mockNationalityRepository := mockRepositories.NewMockINationalityRepository(controller)

	escortHandler := EscortHandler{
		ProfileRepository:               mockProfileRepository,
		ProfileStatusRepository:         mockProfileStatusRepository,
		ProfileStatusCategoryRepository: mockProfileStatusCategoryRepository,
		NationalityRepository:           mockNationalityRepository,
	}

	t.Run(
		"Should interrupt the process when profile status fails for user active account",
		func(t *testing.T) {
			mockProfileStatusRepository.
				EXPECT().
				GetByProfileId(gomock.Any(), gomock.Any()).
				Return(models.ProfileStatus{}, errors.New("dummy error")).
				Times(1)
			mockProfileStatusRepository.
				EXPECT().
				UpdateByProfileId(gomock.Any(), gomock.Any(), gomock.Any()).
				Times(0)
			mockProfileStatusCategoryRepository.
				EXPECT().
				GetOneByName(gomock.Any(), gomock.Any()).
				Times(0)

			escortHandler.HandleEvent(context.Background(), &kafka.Message{
				TopicPartition: kafka.TopicPartition{
					Topic: &config.InitializeKafka().Topics.UserActiveAccount,
				},
				Value: []byte(`{"userId": "638d21e4d9916d56509e66aa"}`),
			})
		})

	t.Run(
		"Should interrupt the process when profile status category fails for user active account",
		func(t *testing.T) {
			mockProfileStatusRepository.
				EXPECT().
				GetByProfileId(gomock.Any(), gomock.Any()).
				Return(models.ProfileStatus{}, nil).
				Times(1)
			mockProfileStatusRepository.
				EXPECT().
				UpdateByProfileId(gomock.Any(), gomock.Any(), gomock.Any()).
				Times(0)
			mockProfileStatusCategoryRepository.
				EXPECT().
				GetOneByName(gomock.Any(), gomock.Any()).
				Return(models.ProfileStatusCategory{}, errors.New("dummy error")).
				Times(1)

			escortHandler.HandleEvent(context.Background(), &kafka.Message{
				TopicPartition: kafka.TopicPartition{
					Topic: &config.InitializeKafka().Topics.UserActiveAccount,
				},
				Value: []byte(`{"userId": "638d21e4d9916d56509e66aa"}`),
			})
		})

	t.Run("Should log error when profile status fails for user active account", func(t *testing.T) {
		mockProfileStatusRepository.
			EXPECT().
			GetByProfileId(gomock.Any(), gomock.Any()).
			Return(models.ProfileStatus{}, nil).
			Times(1)
		mockProfileStatusRepository.
			EXPECT().
			UpdateByProfileId(gomock.Any(), gomock.Any(), gomock.Any()).
			Return(errors.New("dummy error")).
			Times(1)
		mockProfileStatusCategoryRepository.
			EXPECT().
			GetOneByName(gomock.Any(), gomock.Any()).
			Return(models.ProfileStatusCategory{}, nil).
			Times(1)

		escortHandler.HandleEvent(context.Background(), &kafka.Message{
			TopicPartition: kafka.TopicPartition{
				Topic: &config.InitializeKafka().Topics.UserActiveAccount,
			},
			Value: []byte(`{"userId": "638d21e4d9916d56509e66aa"}`),
		})
	})

	t.Run("Should interrupt the process when nationality fails for escort created", func(t *testing.T) {
		mockNationalityRepository.
			EXPECT().
			GetOneByName(gomock.Any(), gomock.Any()).
			Return(models.Nationality{}, errors.New("dummy error")).
			Times(1)
		mockProfileRepository.
			EXPECT().
			Create(gomock.Any(), gomock.Any()).
			Times(0)
		mockProfileStatusCategoryRepository.
			EXPECT().
			GetOneByName(gomock.Any(), gomock.Any()).
			Times(0)
		mockProfileStatusRepository.
			EXPECT().
			Create(gomock.Any(), gomock.Any()).
			Times(0)

		escortHandler.HandleEvent(context.Background(), &kafka.Message{
			TopicPartition: kafka.TopicPartition{
				Topic: &config.InitializeKafka().Topics.EscortCreated,
			},
			Value: []byte(`{"_id": "638d21e4d9916d56509e66aa", "email": "test.user@exampke.com"}`),
		})
	})

	t.Run("Should interrupt the process when profile fails for escort created", func(t *testing.T) {
		mockNationalityRepository.
			EXPECT().
			GetOneByName(gomock.Any(), gomock.Any()).
			Return(models.Nationality{Id: "638d629c7b862642c7c2fb11"}, nil).
			Times(1)
		mockProfileRepository.
			EXPECT().
			Create(gomock.Any(), gomock.Any()).
			Return(errors.New("dummy error")).
			Times(1)
		mockProfileStatusCategoryRepository.
			EXPECT().
			GetOneByName(gomock.Any(), gomock.Any()).
			Times(0)
		mockProfileStatusRepository.
			EXPECT().
			Create(gomock.Any(), gomock.Any()).
			Times(0)

		escortHandler.HandleEvent(context.Background(), &kafka.Message{
			TopicPartition: kafka.TopicPartition{
				Topic: &config.InitializeKafka().Topics.EscortCreated,
			},
			Value: []byte(`{"_id": "638d21e4d9916d56509e66aa", "email": "test.user@exampke.com"}`),
		})
	})

	t.Run(
		"Should interrupt the process when profile status category fails for escort created",
		func(t *testing.T) {
			mockNationalityRepository.
				EXPECT().
				GetOneByName(gomock.Any(), gomock.Any()).
				Return(models.Nationality{Id: "638d629c7b862642c7c2fb11"}, nil).
				Times(1)
			mockProfileRepository.
				EXPECT().
				Create(gomock.Any(), gomock.Any()).
				Return(nil).
				Times(1)
			mockProfileStatusCategoryRepository.
				EXPECT().
				GetOneByName(gomock.Any(), gomock.Any()).
				Return(models.ProfileStatusCategory{}, errors.New("dummy error")).
				Times(1)
			mockProfileStatusRepository.
				EXPECT().
				Create(gomock.Any(), gomock.Any()).
				Times(0)

			escortHandler.HandleEvent(context.Background(), &kafka.Message{
				TopicPartition: kafka.TopicPartition{
					Topic: &config.InitializeKafka().Topics.EscortCreated,
				},
				Value: []byte(`{"_id": "638d21e4d9916d56509e66aa", "email": "test.user@exampke.com"}`),
			})
		})

	t.Run("Should interrupt the process when profile status fails for escort created", func(t *testing.T) {
		mockNationalityRepository.
			EXPECT().
			GetOneByName(gomock.Any(), gomock.Any()).
			Return(models.Nationality{Id: "638d629c7b862642c7c2fb11"}, nil).
			Times(1)
		mockProfileRepository.
			EXPECT().
			Create(gomock.Any(), gomock.Any()).
			Return(nil).
			Times(1)
		mockProfileStatusCategoryRepository.
			EXPECT().
			GetOneByName(gomock.Any(), gomock.Any()).
			Return(models.ProfileStatusCategory{Id: "638d761752bde02d18689c78"}, nil).
			Times(1)
		mockProfileStatusRepository.
			EXPECT().
			Create(gomock.Any(), gomock.Any()).
			Return(errors.New("dummy error")).
			Times(1)

		escortHandler.HandleEvent(context.Background(), &kafka.Message{
			TopicPartition: kafka.TopicPartition{
				Topic: &config.InitializeKafka().Topics.EscortCreated,
			},
			Value: []byte(`{"_id": "638d21e4d9916d56509e66aa", "email": "test.user@exampke.com"}`),
		})
	})
}
