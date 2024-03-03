package user

import (
	"bytes"
	"encoding/json"
	"go_notifier/configs"
	"go_notifier/pkg/database"
	"log"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/suite"
)

var db *database.AppDB

type UserHandlerSuite struct {
	suite.Suite
}

func (suite *UserHandlerSuite) SetupSuite() {
	cfg := configs.InitConfig()
	dbC := cfg.DBConfig

	db = database.InitDB(&dbC)
	err := db.Mysql.Ping()
	if err != nil {
		log.Println("db connection panic")
		panic(err)
	} else {
		log.Println("db connection ok")
	}
	db.RunMigrations("../migrations")
}

func (suite *UserHandlerSuite) TearDownSuite() {
	db.DownMigrations("../migrations")
	db.Mysql.Close()
}

func (suite *UserHandlerSuite) TestCreateUserHandlerSuccess() {
	userService := NewUserService(db)
	userHandler := NewUserHandler(userService)
	ts := httptest.NewServer(http.HandlerFunc(userHandler.CreateUserHandler))
	defer ts.Close()

	requestBody := []byte(`{"UUID": "uuid", "email": "john@example.com", "timezone": "UTC"}`)
	res, err := http.Post(ts.URL+"/user", "application/json", bytes.NewBuffer(requestBody))
	suite.Nil(err)
	defer res.Body.Close()

	suite.Equal(http.StatusCreated, res.StatusCode)
	var response CreateUserResponse
	err = json.NewDecoder(res.Body).Decode(&response)
	suite.Nil(err)
	suite.Equal("uuid", response.UUID)
}

func (suite *UserHandlerSuite) TestCreateUserHandlerFailure() {
	userService := NewUserService(db)
	userHandler := NewUserHandler(userService)
	ts := httptest.NewServer(http.HandlerFunc(userHandler.CreateUserHandler))
	defer ts.Close()

	suite.Run("check validation failure", func() {
		requestBody := []byte(`{"UUID": "", "email": "", "timezone": ""}`)
		res, _ := http.Post(ts.URL+"/user", "application/json", bytes.NewBuffer(requestBody))
		defer res.Body.Close()

		suite.Equal(http.StatusBadRequest, res.StatusCode)
	})

	suite.Run("check duplicate user", func() {
		requestBody := []byte(`{"UUID": "uuidD", "email": "johnD@example.com", "timezone": "UTC"}`)
		res, _ := http.Post(ts.URL+"/user", "application/json", bytes.NewBuffer(requestBody))
		suite.Equal(http.StatusCreated, res.StatusCode)
		res.Body.Close()
		res, _ = http.Post(ts.URL+"/user", "application/json", bytes.NewBuffer(requestBody))
		defer res.Body.Close()

		suite.Equal(http.StatusBadRequest, res.StatusCode)
	})
}

func TestSuite(t *testing.T) {
	suite.Run(t, new(UserHandlerSuite))
}
