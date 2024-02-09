package campaign

import (
	"fmt"
	"go_notifier/internal/common"
	"go_notifier/pkg/database"
	"strings"

	"github.com/google/uuid"
)

type CampaignService struct {
	db *database.AppDB
}

func NewCampaignService(db *database.AppDB) *CampaignService {
	return &CampaignService{
		db: db,
	}
}

func (s *CampaignService) CreateCampaign(dto *common.CampaignRequest) (string, error) {
	newUUID := uuid.New().String()

	insertStatement, err := s.db.Mysql.Prepare("INSERT INTO campaign(uuid, name, message, time, days_of_week) VALUES (?, ?, ?, ?, ?)")
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
