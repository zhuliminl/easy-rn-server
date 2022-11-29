package service

import (
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
	return u.userRepository.CreateUser(userDto)
}

func (u userService) GetUserByUserId(userId string) (dto.User, error) {
	userEntity, err := u.userRepository.GetUserById(userId)
	if err != nil {
		return dto.User{}, err
	}
	user := dto.User{
		Username: userEntity.Username,
		UserId:   userEntity.ID,
	}
	return user, nil
}

func NewUserService(userRepo repository.UserRepository) UserService {
	return &userService{
		userRepository: userRepo,
	}
}

func MapEntityUserToUser(user entity.User) dto.User {
	return dto.User{
		UserId:         user.ID,
		Username:       user.Username,
		Email:          user.Email,
		Phone:          user.Phone,
		WechatNickname: user.WechatNickname,
	}
}
