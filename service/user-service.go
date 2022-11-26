package service

import (
	"github.com/zhuliminl/easyrn-server/db"
	"github.com/zhuliminl/easyrn-server/dto"
	"github.com/zhuliminl/easyrn-server/entity"
	"github.com/zhuliminl/easyrn-server/repository"
)

type UserService interface {
	GetUserByUserId(userId string) (dto.User, error)
	CreateUser(user dto.User) error
}

type userService struct {
	userRepository repository.UserRepository
}

func (u userService) CreateUser(userDto dto.User) error {
	var user entity.User
	user.Name = userDto.Name
	//user.Email = "zhuliminl@gmial.com"
	return u.userRepository.CreateUser(user)
}

func (u userService) GetUserByUserId(userId string) (dto.User, error) {
	u.CreateUser(dto.User{})

	var user dto.User
	user.Name = "saul"
	if err := db.DB.Where("name = ?", userId).First(&user).Error; err != nil {
		return user, err
	}

	return user, nil
}

func NewUserService(userRepo repository.UserRepository) UserService {
	return &userService{
		userRepository: userRepo,
	}
}
