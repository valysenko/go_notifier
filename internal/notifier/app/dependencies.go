package app

import (
	"go_notifier/internal/campaign"
	"os"

	log "github.com/sirupsen/logrus"
)

func NewLogger() log.FieldLogger {
	logger := log.New()
	logger.SetOutput(os.Stdout)
	logger.SetFormatter(&log.JSONFormatter{DataKey: "fields"})

	return logger
}

func NewCampaignRepository(app *NotifierApp) *campaign.MysqlCampaignRepository {
	return campaign.NewMysqlCampaignRepository(app.mysql)
}

func NewCampaignService(app *NotifierApp) *campaign.CampaignService {
	return campaign.NewCampaignService(app.mysql, NewCampaignRepository(app), app.logger)
}
