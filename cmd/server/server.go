package main

import (
	"context"
	"go_notifier/configs"
	"go_notifier/internal/http/app"
)

func main() {
	ctx := context.Background()
	cfg := configs.InitConfig()
	app := app.NewServerApp(ctx, cfg)
	defer app.Close(ctx)
	app.Run(ctx)
}
