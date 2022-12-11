package service

import (
	"github.com/zhuliminl/easyrn-server/repository"
)

type SettingService interface {
}

type settingService struct {
	userRepository repository.UserRepository
}

func NewSettingService(userRepo repository.UserRepository) SettingService {
	return &settingService{
		userRepository: userRepo,
	}
}
