package campaign

import (
	"fmt"
	"go_notifier/internal/common"
	"go_notifier/pkg/database"
	"strconv"
	"strings"
	"time"

	"github.com/google/uuid"
	log "github.com/sirupsen/logrus"
)

type CampaignService struct {
	db         *database.AppDB
	repository CampaignRepository
	logger     log.FieldLogger
}

type CampaignRepository interface {
	GetScheduledNotifications(day, currentTime string) ([]*ScheduledNotification, error)
}

func NewCampaignService(db *database.AppDB, repo CampaignRepository, logger log.FieldLogger) *CampaignService {
	return &CampaignService{
		db:         db,
		repository: repo,
		logger:     logger,
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

func (s *CampaignService) SendScheduledNotifications() error {
	currentTime := time.Now().UTC().Truncate(time.Minute)
	currentDay := int(currentTime.Weekday())
	currentTimeString := currentTime.Format(time.TimeOnly)
	s.logger.Infof("sending notifications on %d weekday at %s", currentDay, currentTimeString)

	res, err := s.repository.GetScheduledNotifications(strconv.Itoa(currentDay), currentTimeString)
	if err != nil {
		return err
	}
	for _, r := range res {
		fmt.Println(r)
	}

	return nil
}
