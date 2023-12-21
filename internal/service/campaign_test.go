package service

import (
	"errors"
	"go_notifier/internal/dto"
	"go_notifier/pkg/database"
	"strings"
	"testing"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/stretchr/testify/assert"
)

func TestCreateCampaign(t *testing.T) {
	db, mock := database.InitMockDB(t)
	defer db.Close()

	dto := &dto.Campaign{
		Name:       "MyCampaign",
		Message:    "Hello, world!",
		Time:       "12:00 PM",
		DaysOfWeek: []string{"1", "3"},
	}

	t.Run("create campaign success", func(t *testing.T) {
		mock.ExpectPrepare("INSERT INTO campaign").
			ExpectExec().
			WithArgs(sqlmock.AnyArg(), dto.Name, dto.Message, dto.Time, strings.Join(dto.DaysOfWeek, ",")).
			WillReturnResult(sqlmock.NewResult(1, 1))

		uuid, err := CreateCampaign(dto)
		assert.Nil(t, err)
		assert.Len(t, uuid, 36)
	})

	t.Run("create campaign prepare failure", func(t *testing.T) {
		expectedError := errors.New("error")
		mock.ExpectPrepare("INSERT INTO campaign").WillReturnError(expectedError)
		uuid, err := CreateCampaign(dto)
		assert.NotNil(t, err)
		assert.Equal(t, expectedError, err)
		assert.Equal(t, "", uuid)
	})

	t.Run("create campaign exec failure", func(t *testing.T) {
		expectedError := errors.New("error")
		mock.ExpectPrepare("INSERT INTO campaign").
			ExpectExec().
			WithArgs(sqlmock.AnyArg(), dto.Name, dto.Message, dto.Time, strings.Join(dto.DaysOfWeek, ",")).
			WillReturnError(expectedError)
		uuid, err := CreateCampaign(dto)
		assert.NotNil(t, err)
		assert.Equal(t, expectedError, err)
		assert.Equal(t, "", uuid)
	})
}
