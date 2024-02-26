package user_app

import (
	"go_notifier/internal/common"
	"go_notifier/pkg/database"
)

// go:generate go run github.com/vektra/mockery/v2@v2.20.0 --name=UserRepository --case snake
type UserRepository interface {
	GetUserIDByUUID(uuid string) (int64, error)
}

type UserAppService struct {
	db             *database.AppDB
	userRepository UserRepository
}

func NewUserAppService(db *database.AppDB, userRepo UserRepository) *UserAppService {
	return &UserAppService{
		db:             db,
		userRepository: userRepo,
	}
}

func (s *UserAppService) CreateUserApp(request *common.UserAppRequest) (int64, error) {
	userId, err := s.userRepository.GetUserIDByUUID(request.UserUUID)
	if err != nil {
		return 0, err
	}

	insertStatement, err := s.db.Mysql.Prepare("INSERT INTO user_app(identifier, type, user_id) VALUES (?, ?, ?)")
	if err != nil {
		return 0, err
	}
	defer insertStatement.Close()

	res, err := insertStatement.Exec(request.Identifier, request.Type, userId)
	if err != nil {
		return 0, err
	}

	lastId, err := res.LastInsertId()
	return lastId, err
}
