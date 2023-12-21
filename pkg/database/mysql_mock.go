package database

import (
	"database/sql"
	"testing"

	"github.com/DATA-DOG/go-sqlmock"
)

func InitMockDB(t *testing.T) (*sql.DB, sqlmock.Sqlmock) {
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("An error occurred while creating mock: %s", err)
	}
	DB = &AppDB{
		Mysql: db,
	}

	return db, mock
}
