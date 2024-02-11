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

type ScheduledNotification struct {
	CampaignUuid string
	CampaignName string
	Message      string
	UserUuid     string
	DeviceToken  string
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

func (repo *MysqlCampaignRepository) GetScheduledNotifications(day, currentTime string) ([]*ScheduledNotification, error) {
	var notifications []*ScheduledNotification

	query := `SELECT campaign.uuid, campaign.name, campaign.message, u.uuid as user_uuid, d.token
		FROM campaign
		LEFT JOIN user_campaign ON user_campaign.campaign_id = campaign.id
		LEFT JOIN user u ON user_campaign.user_id = u.id
		LEFT JOIN device d ON u.id = d.user_id
		WHERE campaign.is_active = true
	  	AND FIND_IN_SET(?, days_of_week)
	 	AND user_campaign.time = ?;`

	stmt, err := repo.db.Mysql.Prepare(query)
	if err != nil {
		return nil, err
	}
	defer stmt.Close()

	rows, err := stmt.Query(day, currentTime)
	if err != nil {
		panic(err.Error())
	}
	defer rows.Close()

	for rows.Next() {
		var notification ScheduledNotification
		err := rows.Scan(&notification.CampaignUuid, &notification.CampaignName, &notification.Message, &notification.UserUuid, &notification.DeviceToken)
		if err != nil {
			return nil, err
		}
		notifications = append(notifications, &notification)
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}

	return notifications, nil
}
