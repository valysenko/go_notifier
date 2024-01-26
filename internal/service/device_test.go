package service

import (
	"go_notifier/internal/dto"
	"go_notifier/internal/service/mocks"
	"go_notifier/pkg/database"
	"testing"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/stretchr/testify/assert"
)

func TestCreateDevice(t *testing.T) {
	db, mock := database.InitMockDB(t)
	defer db.Close()

	mockUserRepo := mocks.NewUserRepository(t)
	SetUserRepository(mockUserRepo)

	dto := &dto.Device{
		Token:    "token",
		UserUUID: "uuid",
	}

	t.Run("create device success", func(t *testing.T) {
		var expectedId int64 = 4
		mockUserRepo.On("GetUserIDByUUID", dto.UserUUID).Once().Return(expectedId, nil)
		mock.ExpectPrepare("INSERT INTO device").
			ExpectExec().
			WithArgs(dto.Token, expectedId).
			WillReturnResult(sqlmock.NewResult(expectedId, 1))

		id, err := CreateDevice(dto)
		assert.Nil(t, err)
		assert.Equal(t, expectedId, id)
	})
}
