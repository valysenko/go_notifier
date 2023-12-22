package repository

import (
	"database/sql"
	"fmt"
	"go_notifier/pkg/database"
)

type CampaignIdTime struct {
	ID   int
	Time string
}

//go:generate go run github.com/vektra/mockery/v2@v2.20.0 --name=CampaignRepository --case snake
type CampaignRepository interface {
	GetCampgignIdAndTimeByUUID(uuid string) (*CampaignIdTime, error)
}

type CampaignRepositoryImpl struct {
}

func (repo *CampaignRepositoryImpl) GetCampgignIdAndTimeByUUID(uuid string) (*CampaignIdTime, error) {
	var campaign CampaignIdTime
	query := "SELECT id, time FROM campaign WHERE uuid = ?"

	err := database.DB.Mysql.QueryRow(query, uuid).Scan(&campaign.ID, &campaign.Time)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, fmt.Errorf("No campaign found with the given UUID=%s", uuid)
		} else {
			return nil, err
		}
	}

	return &campaign, nil
}
