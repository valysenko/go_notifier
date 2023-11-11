package device

import (
	"go_notifier/internal/db/repository"
	"go_notifier/internal/dto"
	"go_notifier/pkg/database"
)

func CreateDevice(dto *dto.Device) (int64, error) {
	userId, err := repository.GetUserIDByUUID(dto.UserUUID)
	if err != nil {
		return 0, err
	}

	insertStatement, err := database.DB.Mysql.Prepare("INSERT INTO device(token, user_id) VALUES (?, ?)")
	if err != nil {
		return 0, err
	}
	defer insertStatement.Close()

	res, err := insertStatement.Exec(dto.Token, userId)
	if err != nil {
		return 0, err
	}

	lastId, err := res.LastInsertId()
	return lastId, err
}
