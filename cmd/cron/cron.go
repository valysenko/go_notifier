package main

import (
	"context"
	"go_notifier/configs"
	"go_notifier/internal/app/notifier/app"

	_ "github.com/go-sql-driver/mysql"
)

func main() {
	ctx := context.Background()
	cfg := configs.InitConfig()
	app := app.NewNotifierApp(ctx, cfg)
	defer app.Close(ctx)
	app.Run(ctx)
}
