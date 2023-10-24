package database

import (
	"database/sql"
	"go_notifier/configs"
)

type DB struct {
	Mysql *sql.DB
}

func InitDB(dbConfig *configs.DBConfig) *DB {
	db, err := sql.Open("mysql", dbConfig.ProvideDSN())
	if err != nil {
		panic(err)
	}
	db.SetConnMaxLifetime(dbConfig.ConnMaxLifetime)
	db.SetMaxOpenConns(dbConfig.MaxOpenConns)
	db.SetMaxIdleConns(dbConfig.MaxIdleConns)

	return &DB{
		Mysql: db,
	}
}
