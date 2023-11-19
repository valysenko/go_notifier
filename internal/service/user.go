package service

import (
	"go_notifier/internal/dto"
	"go_notifier/pkg/database"
)

func CreateUser(dto *dto.User) (int64, error) {
	insertStatement, err := database.DB.Mysql.Prepare("INSERT INTO user(uuid, email, timezone) VALUES (?, ?, ?)")
	if err != nil {
		return 0, err
	}
	defer insertStatement.Close()

	res, err := insertStatement.Exec(dto.UUID, dto.Email, dto.Timezone)
	if err != nil {
		return 0, err
	}

	lastId, err := res.LastInsertId()
	return lastId, err
}
