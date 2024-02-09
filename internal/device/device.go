package device

import (
	"go_notifier/internal/common"
	"go_notifier/pkg/database"
)

// go:generate go run github.com/vektra/mockery/v2@v2.20.0 --name=UserRepository --case snake
type UserRepository interface {
	GetUserIDByUUID(uuid string) (int64, error)
}

type DeviceService struct {
	db             *database.AppDB
	userRepository UserRepository
}

func NewDeviceService(db *database.AppDB, userRepo UserRepository) *DeviceService {
	return &DeviceService{
		db:             db,
		userRepository: userRepo,
	}
}

func (s *DeviceService) CreateDevice(request *common.DeviceRequest) (int64, error) {
	userId, err := s.userRepository.GetUserIDByUUID(request.UserUUID)
	if err != nil {
		return 0, err
	}

	insertStatement, err := s.db.Mysql.Prepare("INSERT INTO device(token, user_id) VALUES (?, ?)")
	if err != nil {
		return 0, err
	}
	defer insertStatement.Close()

	res, err := insertStatement.Exec(request.Token, userId)
	if err != nil {
		return 0, err
	}

	lastId, err := res.LastInsertId()
	return lastId, err
}
