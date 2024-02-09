package database

import (
	"testing"

	"github.com/DATA-DOG/go-sqlmock"
)

func InitMockDB(t *testing.T) (*AppDB, sqlmock.Sqlmock) {
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("An error occurred while creating mock: %s", err)
	}
	return &AppDB{
		Mysql: db,
	}, mock
}
