package service

import (
	"github.com/zhuliminl/easyrn-server/dto"
	"github.com/zhuliminl/easyrn-server/repository"
)

type UserService interface {
	GetUserById(id string) (dto.User, error)
}

type userService struct {
	userRepository repository.UserRepository
}

func (u userService) GetUserById(id string) (dto.User, error) {
	var user dto.User
	user.Name = "saul"
	if err := 
	return user, nil
}

func NewUserService(userRepo repository.UserRepository) UserService {
	return &userService{
		userRepository: userRepo,
	}
}
