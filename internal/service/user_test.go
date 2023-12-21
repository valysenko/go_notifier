package service

import (
	"errors"
	"go_notifier/internal/dto"
	"go_notifier/pkg/database"
	"testing"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/stretchr/testify/assert"
)

func TestCreateUser(t *testing.T) {
	db, mock := database.InitMockDB(t)
	defer db.Close()

	dto := &dto.User{
		UUID:     "uuid",
		Email:    "email",
		Timezone: "Europe/Kyiv",
	}

	t.Run("create user success", func(t *testing.T) {
		var expectedId int64 = 3
		mock.ExpectPrepare("INSERT INTO user").
			ExpectExec().
			WithArgs(dto.UUID, dto.Email, dto.Timezone).
			WillReturnResult(sqlmock.NewResult(expectedId, 1))

		id, err := CreateUser(dto)
		assert.Nil(t, err)
		assert.Equal(t, expectedId, id)
	})

	t.Run("create user exec failure", func(t *testing.T) {
		expectedError := errors.New("error")

		mock.ExpectPrepare("INSERT INTO user").
			ExpectExec().
			WithArgs(dto.UUID, dto.Email, dto.Timezone).WillReturnError(expectedError)

		id, err := CreateUser(dto)
		assert.NotNil(t, err)
		assert.Equal(t, int64(0), id)

	})
}
