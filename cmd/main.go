package main

import (
	"go_notifier/configs"
	"go_notifier/pkg/database"
	server "go_notifier/pkg/http"
	"log"

	_ "github.com/go-sql-driver/mysql"
)

// go run cmd/main.go
func main() {
	appConfig := configs.InitConfig()

	// mysql
	db := database.InitDB(&appConfig.DBConfig)
	defer db.Mysql.Close()

	err := db.Mysql.Ping()
	if err != nil {
		log.Println("db connection panic")
		panic(err)
	} else {
		log.Println("db connection ok")
	}

	db.RunMigrations()

	// http server
	httpServer := server.InitServer(&appConfig.HttpServerConfig)
	if err := httpServer.Start(); err != nil {
		panic(err)
	}
}
