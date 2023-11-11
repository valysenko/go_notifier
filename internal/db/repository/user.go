package repository

import (
	"database/sql"
	"fmt"
	"go_notifier/pkg/database"
)

func GetUserIDByUUID(uuid string) (int64, error) {
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
