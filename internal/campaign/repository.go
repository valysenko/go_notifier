package campaign

import (
	"database/sql"
	"fmt"
	"go_notifier/pkg/database"
)

type CampaignIdTime struct {
	ID   int
	Time string
}

type MysqlCampaignRepository struct {
	db *database.AppDB
}

func NewMysqlCampaignRepository(db *database.AppDB) *MysqlCampaignRepository {
	return &MysqlCampaignRepository{
		db: db,
	}
}

func (repo *MysqlCampaignRepository) GetCampgignIdAndTimeByUUID(uuid string) (*CampaignIdTime, error) {
	var campaign CampaignIdTime
	query := "SELECT id, time FROM campaign WHERE uuid = ?"

	err := repo.db.Mysql.QueryRow(query, uuid).Scan(&campaign.ID, &campaign.Time)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, fmt.Errorf("No campaign found with the given UUID=%s", uuid)
		} else {
			return nil, err
		}
	}

	return &campaign, nil
}
