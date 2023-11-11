package database

import (
	"database/sql"
	"go_notifier/configs"

	"github.com/pressly/goose"
)

var DB *AppDB

type AppDB struct {
	Mysql *sql.DB
}

func InitDB(dbConfig *configs.DBConfig) *AppDB {
	db, err := sql.Open("mysql", dbConfig.ProvideDSN())
	if err != nil {
		panic(err)
	}
	db.SetConnMaxLifetime(dbConfig.ConnMaxLifetime)
	db.SetMaxOpenConns(dbConfig.MaxOpenConns)
	db.SetMaxIdleConns(dbConfig.MaxIdleConns)

	DB = &AppDB{
		Mysql: db,
	}

	return DB
}

// goose create create_user_campaign_table sql
// goose create create_user_campaign_table go
// goose mysql "admin:go_notifier@tcp(127.0.0.1:23306)/go_notifier" up
func (db *AppDB) RunMigrations() {
	goose.SetDialect("mysql")
	err := goose.Up(db.Mysql, "internal/db/migrations")
	if err != nil {
		panic(err)
	}
}
