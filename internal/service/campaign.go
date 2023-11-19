package service

import (
	"fmt"
	"go_notifier/internal/dto"
	"go_notifier/pkg/database"
	"strings"

	"github.com/google/uuid"
)

func CreateCampaign(dto *dto.Campaign) (string, error) {
	newUUID := uuid.New().String()

	insertStatement, err := database.DB.Mysql.Prepare("INSERT INTO campaign(uuid, name, message, time, days_of_week) VALUES (?, ?, ?, ?, ?)")
	if err != nil {
		fmt.Println(err)
		return "", err
	}
	defer insertStatement.Close()

	res, err := insertStatement.Exec(newUUID, dto.Name, dto.Message, dto.Time, strings.Join(dto.DaysOfWeek, ","))
	if err != nil {
		return "", err
	}

	_, err = res.LastInsertId()
	return newUUID, err
}
