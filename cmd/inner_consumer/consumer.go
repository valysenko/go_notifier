package main

import (
	"context"
	"go_notifier/configs"
	"go_notifier/internal/app/consumer/inner/app"

	_ "github.com/go-sql-driver/mysql"
)

func main() {
	ctx := context.Background()
	cfg := configs.InitConfig()
	app := app.NewConsumerApp(ctx, cfg)
	defer app.Close(ctx)
	err := app.Run(ctx)
	if err != nil {
		app.LogError(err)
	}
}
