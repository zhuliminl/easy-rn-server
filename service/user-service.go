package service

import (
	"github.com/zhuliminl/easyrn-server/entity"
	"github.com/zhuliminl/easyrn-server/repository"
)

type UserService interface {
	GetUserByUserId(userId string) (entity.User, error)
	CreateUser(user entity.User) error
}

type userService struct {
	userRepository repository.UserRepository
}

func (u userService) CreateUser(user entity.User) error {
	return u.userRepository.CreateUser(user)
}

func (u userService) GetUserByUserId(userId string) (entity.User, error) {
	return u.userRepository.GetUserById(userId)
}

func NewUserService(userRepo repository.UserRepository) UserService {
	return &userService{
		userRepository: userRepo,
	}
}
