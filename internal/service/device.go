package service

import (
	"go_notifier/internal/dto"
	"go_notifier/pkg/database"
)

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

func (s *DeviceService) CreateDevice(dto *dto.Device) (int64, error) {
	userId, err := s.userRepository.GetUserIDByUUID(dto.UserUUID)
	if err != nil {
		return 0, err
	}

	insertStatement, err := s.db.Mysql.Prepare("INSERT INTO device(token, user_id) VALUES (?, ?)")
	if err != nil {
		return 0, err
	}
	defer insertStatement.Close()

	res, err := insertStatement.Exec(dto.Token, userId)
	if err != nil {
		return 0, err
	}

	lastId, err := res.LastInsertId()
	return lastId, err
}
