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
	publisher  Publisher
}

// go:generate go run github.com/vektra/mockery/v2@v2.20.0 --name=CampaignRepository --case snake
type CampaignRepository interface {
	GetScheduledNotifications(day, currentTime string) ([]*ScheduledNotification, error)
}

// moved to publisher folder
// go:generate go run github.com/vektra/mockery/v2@v2.20.0 --name=Publisher --case snake
type Publisher interface {
	Publish(queueName string, event []byte) error
}

func NewCampaignService(db *database.AppDB, repo CampaignRepository, logger log.FieldLogger, publisher Publisher) *CampaignService {
	return &CampaignService{
		db:         db,
		repository: repo,
		logger:     logger,
		publisher:  publisher,
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
		err := s.publishScheduledNotificationEvent(r)
		if err != nil {
			s.logger.WithFields(log.Fields{
				"notification": r,
			}).Error("failed to publish scheduled notification")
		}
	}

	return nil
}

func (s *CampaignService) publishScheduledNotificationEvent(n *ScheduledNotification) error {
	event, err := common.NewScheduledNotification(n.CampaignName, n.Message, n.UserUuid, n.AppIdentifier, n.AppType)
	if err != nil {
		return fmt.Errorf("failed to create zendesk ticket comment event %w", err)
	}

	return s.publisher.Publish(common.ScheduledNotificationsQueue, event)
}
