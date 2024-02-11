package app

import (
	"context"
	"go_notifier/configs"
	"go_notifier/pkg/database"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/go-co-op/gocron/v2"
	log "github.com/sirupsen/logrus"
)

type NotifierApp struct {
	cfg    *configs.AppConfig
	logger log.FieldLogger
	mysql  *database.AppDB
}

func NewNotifierApp(ctx context.Context, cfg *configs.AppConfig) *NotifierApp {
	logger := NewLogger()
	appConfig := configs.InitConfig()

	db := database.InitDB(&appConfig.DBConfig)
	err := db.Mysql.Ping()
	if err != nil {
		panic(err)
	}

	return &NotifierApp{
		cfg:    appConfig,
		logger: logger,
		mysql:  db,
	}
}

func (app *NotifierApp) Run(ctx context.Context) {
	campaignService := NewCampaignService(app)

	s, err := gocron.NewScheduler()
	if err != nil {
		app.logger.Errorf("failed to start notifier because of %s", err.Error())
		os.Exit(1)
		return
	}

	// configure task to run every minute
	_, err = s.NewJob(
		gocron.DurationJob(
			1*time.Minute,
		),
		gocron.NewTask(func() {
			if err := campaignService.SendScheduledNotifications(); err != nil {
				app.logger.Errorf("notifier failure because of %s", err.Error())
				os.Exit(1)
			}
		}),
	)
	if err != nil {
		app.logger.Errorf("notifier failure because of %s", err.Error())
		os.Exit(1)
	}

	// start scheduler
	s.Start()

	stop := make(chan os.Signal, 1)
	signal.Notify(stop, syscall.SIGINT, syscall.SIGTERM)
	<-stop

	err = s.Shutdown()
	if err != nil {
		app.logger.Errorf("failed to shutdown notifier because of %s", err.Error())
		os.Exit(1)
	}
}

func (app *NotifierApp) Close(ctx context.Context) {
	app.mysql.Mysql.Close()
}
