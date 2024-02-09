package user

import (
	"fmt"
	"go_notifier/internal/common"
	"go_notifier/pkg/database"
)

type UserService struct {
	db *database.AppDB
}

func NewUserService(db *database.AppDB) *UserService {
	return &UserService{
		db: db,
	}
}

func (s *UserService) CreateUser(request *common.UserRequest) (int64, error) {
	insertStatement, err := s.db.Mysql.Prepare("INSERT INTO user(uuid, email, timezone) VALUES (?, ?, ?)")
	if err != nil {
		return 0, err
	}
	defer insertStatement.Close()

	res, err := insertStatement.Exec(request.UUID, request.Email, request.Timezone)
	if err != nil {
		return 0, err
	}

	lastId, err := res.LastInsertId()
	fmt.Println(err)
	return lastId, err
}
