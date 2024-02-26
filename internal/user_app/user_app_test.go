package user_app

import (
	"go_notifier/internal/common"
	"go_notifier/internal/user_app/mocks"
	"go_notifier/pkg/database"
	"testing"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/stretchr/testify/assert"
)

func TestCreateDevice(t *testing.T) {
	db, mock := database.InitMockDB(t)
	defer db.Mysql.Close()
	mockUserRepo := mocks.NewUserRepository(t)
	s := NewUserAppService(db, mockUserRepo)

	dto := &common.UserAppRequest{
		Identifier: "id",
		Type:       common.FIREBASE_APP,
		UserUUID:   "uuid",
	}

	t.Run("create user app success", func(t *testing.T) {
		var expectedId int64 = 4
		mockUserRepo.On("GetUserIDByUUID", dto.UserUUID).Once().Return(expectedId, nil)
		mock.ExpectPrepare("INSERT INTO user_app").
			ExpectExec().
			WithArgs(dto.Identifier, dto.Type, expectedId).
			WillReturnResult(sqlmock.NewResult(expectedId, 1))

		id, err := s.CreateUserApp(dto)
		assert.Nil(t, err)
		assert.Equal(t, expectedId, id)
	})
}
