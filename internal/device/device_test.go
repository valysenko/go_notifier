package device

import (
	"go_notifier/internal/common"
	"go_notifier/internal/device/mocks"
	"go_notifier/pkg/database"
	"testing"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/stretchr/testify/assert"
)

func TestCreateDevice(t *testing.T) {
	db, mock := database.InitMockDB(t)
	defer db.Mysql.Close()
	mockUserRepo := mocks.NewUserRepository(t)
	s := NewDeviceService(db, mockUserRepo)

	dto := &common.DeviceRequest{
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

		id, err := s.CreateDevice(dto)
		assert.Nil(t, err)
		assert.Equal(t, expectedId, id)
	})
}
