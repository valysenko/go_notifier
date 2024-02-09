package app

import (
	"go_notifier/internal/db/repository"
	"go_notifier/internal/http/handlers"
	"go_notifier/internal/http/server"
	"go_notifier/internal/service"
	"os"

	log "github.com/sirupsen/logrus"
)

func NewLogger() log.FieldLogger {
	logger := log.New()
	logger.SetOutput(os.Stdout)
	logger.SetFormatter(&log.JSONFormatter{DataKey: "fields"})

	return logger
}

// repository
func NewUserRepository(app *ServerApp) *repository.MysqlUserRepository {
	return repository.NewMysqlUserRepository(app.mysql)
}

func NewCampaignRepository(app *ServerApp) *repository.MysqlCampaignRepository {
	return repository.NewMysqlCampaignRepository(app.mysql)
}

// service
func NewDeviceService(app *ServerApp) *service.DeviceService {
	return service.NewDeviceService(app.mysql, NewUserRepository(app))
}

func NewCampaignService(app *ServerApp) *service.CampaignService {
	return service.NewCampaignService(app.mysql)
}

func NewUserService(app *ServerApp) *service.UserService {
	return service.NewUserService(app.mysql)
}

func NewUserCampaignService(app *ServerApp) *service.UserCampaignService {
	return service.NewUserCampaignService(app.mysql, NewUserRepository(app), NewCampaignRepository(app))
}

// http handler
func NewDeviceHandler(app *ServerApp) *handlers.DeviceHandler {
	return handlers.NewDeviceHandler(NewDeviceService(app))
}

func NewCampaignHandler(app *ServerApp) *handlers.CampaignHandler {
	return handlers.NewCampaignHandler(NewCampaignService(app))
}

func NewUserHandler(app *ServerApp) *handlers.UserHandler {
	return handlers.NewUserHandler(NewUserService(app))
}

func NewUserCampaignHandler(app *ServerApp) *handlers.UserCampaignHandler {
	return handlers.NewUserCampaignHandler(NewUserCampaignService(app))
}

// http server
func NewServer(app *ServerApp) *server.HttpServer {
	dh := NewDeviceHandler(app)
	ch := NewCampaignHandler(app)
	uh := NewUserHandler(app)
	uch := NewUserCampaignHandler(app)
	return server.InitServer(&app.cfg.HttpServerConfig, app.logger, dh, ch, uh, uch)
}
