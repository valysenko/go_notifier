package main

import (
	"fmt"
	"go_notifier/configs"
	"go_notifier/pkg/database"
	"time"

	_ "github.com/go-sql-driver/mysql"
)

func main() {
	appConfig := configs.InitConfig()
	db := database.InitDB(&appConfig.DBConfig)
	defer db.Mysql.Close()

	err := db.Mysql.Ping()
	if err != nil {
		fmt.Println("poanic 1")
		panic(err)
	}
	if err == nil {
		fmt.Println("no panic 1")
	}

	db.RunMigrations()

	time.Sleep(time.Minute * 20)
}
