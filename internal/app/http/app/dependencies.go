package app

import (
	"go_notifier/internal/app/http/server"
	"go_notifier/internal/campaign"
	"go_notifier/internal/user"
	"go_notifier/internal/user_app"
	"go_notifier/internal/user_campaign"
	"os"

	log "github.com/sirupsen/logrus"
)

func NewLogger() log.FieldLogger {
	logger := log.New()
	logger.SetOutput(os.Stdout)
	// logger.SetFormatter(&log.JSONFormatter{DataKey: "fields"})

	return logger
}

// repository
func NewUserRepository(app *ServerApp) *user.MysqlUserRepository {
	return user.NewMysqlUserRepository(app.mysql)
}

func NewCampaignRepository(app *ServerApp) *campaign.MysqlCampaignRepository {
	return campaign.NewMysqlCampaignRepository(app.mysql)
}

// service
func NewUserAppService(app *ServerApp) *user_app.UserAppService {
	return user_app.NewUserAppService(app.mysql, NewUserRepository(app))
}

func NewCampaignService(app *ServerApp) *campaign.CampaignService {
	return campaign.NewCampaignService(app.mysql, NewCampaignRepository(app), app.logger, nil)
}

func NewUserService(app *ServerApp) *user.UserService {
	return user.NewUserService(app.mysql)
}

func NewUserCampaignService(app *ServerApp) *user_campaign.UserCampaignService {
	return user_campaign.NewUserCampaignService(app.mysql, NewUserRepository(app), NewCampaignRepository(app))
}

// http handler
func NewUserAppHandler(app *ServerApp) *user_app.UserAppHandler {
	return user_app.NewUserAppHandler(NewUserAppService(app))
}

func NewCampaignHandler(app *ServerApp) *campaign.CampaignHandler {
	return campaign.NewCampaignHandler(NewCampaignService(app))
}

func NewUserHandler(app *ServerApp) *user.UserHandler {
	return user.NewUserHandler(NewUserService(app))
}

func NewUserCampaignHandler(app *ServerApp) *user_campaign.UserCampaignHandler {
	return user_campaign.NewUserCampaignHandler(NewUserCampaignService(app))
}

// http server
func NewServer(app *ServerApp) *server.HttpServer {
	dh := NewUserAppHandler(app)
	ch := NewCampaignHandler(app)
	uh := NewUserHandler(app)
	uch := NewUserCampaignHandler(app)
	return server.InitServer(&app.cfg.HttpServerConfig, app.logger, dh, ch, uh, uch)
}
