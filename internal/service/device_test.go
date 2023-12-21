package service

import (
	"go_notifier/internal/dto"
	"go_notifier/pkg/database"
	"testing"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/stretchr/testify/assert"
)

func TestCreateDevice(t *testing.T) {
	db, mock := database.InitMockDB(t)
	defer db.Close()

	dto := &dto.Device{
		Token:    "token",
		UserUUID: "uuid",
	}

	t.Run("create device success", func(t *testing.T) {
		var expectedId int64 = 4
		// not OK to mock inner's function call - better to have an repo interface to moch a whole function
		mock.ExpectQuery("SELECT id FROM user WHERE uuid = ?").WithArgs(dto.UserUUID).WillReturnRows(sqlmock.NewRows([]string{"id"}).AddRow(expectedId))
		mock.ExpectPrepare("INSERT INTO device").
			ExpectExec().
			WithArgs(dto.Token, expectedId).
			WillReturnResult(sqlmock.NewResult(expectedId, 1))

		id, err := CreateDevice(dto)
		assert.Nil(t, err)
		assert.Equal(t, expectedId, id)
	})
}
