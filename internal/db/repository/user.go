package repository

import (
	"database/sql"
	"fmt"
	"go_notifier/pkg/database"
)

type UserIdTimezone struct {
	ID       int
	Timezone string
}

//go:generate go run github.com/vektra/mockery/v2@v2.20.0 --name=UserRepository --case snake
type UserRepository interface {
	GetUserIDByUUID(uuid string) (int64, error)
	GetUserIDAndTimezoneByUUID(uuid string) (*UserIdTimezone, error)
}

type UserRepositoryImpl struct {
}

func (repo *UserRepositoryImpl) GetUserIDByUUID(uuid string) (int64, error) {
	var userID int64
	query := "SELECT id FROM user WHERE uuid = ?"

	err := database.DB.Mysql.QueryRow(query, uuid).Scan(&userID)
	if err != nil {
		if err == sql.ErrNoRows {
			return 0, fmt.Errorf("No user found with the given UUID=%s", uuid)
		} else {
			return 0, err
		}
	}

	return userID, nil
}

func (repo *UserRepositoryImpl) GetUserIDAndTimezoneByUUID(uuid string) (*UserIdTimezone, error) {
	var user UserIdTimezone
	query := "SELECT id, timezone FROM user WHERE uuid = ?"

	err := database.DB.Mysql.QueryRow(query, uuid).Scan(&user.ID, &user.Timezone)

	if err != nil {
		if err == sql.ErrNoRows {
			return nil, fmt.Errorf("No user found with the given UUID=%s", uuid)
		} else {
			return nil, err
		}
	}

	return &user, nil
}
