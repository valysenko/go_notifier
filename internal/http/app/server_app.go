package app

import (
	"context"
	"go_notifier/configs"
	"go_notifier/pkg/database"

	log "github.com/sirupsen/logrus"
)

type ServerApp struct {
	cfg    *configs.AppConfig
	logger log.FieldLogger
	mysql  *database.AppDB
}

func NewServerApp(ctx context.Context, cfg *configs.AppConfig) *ServerApp {
	logger := NewLogger()
	appConfig := configs.InitConfig()

	db := database.InitDB(&appConfig.DBConfig)
	err := db.Mysql.Ping()
	if err != nil {
		panic(err)
	}
	db.RunMigrations("internal/db/migrations")

	return &ServerApp{
		cfg:    appConfig,
		logger: logger,
		mysql:  db,
	}
}

func (app *ServerApp) Run(ctx context.Context) {
	httpServer := NewServer(app)
	if err := httpServer.Start(); err != nil {
		panic(err)
	}
}

func (app *ServerApp) Close(ctx context.Context) {
	app.mysql.Mysql.Close()
}
