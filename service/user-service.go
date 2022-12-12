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

func (u userService) CreateUser(userDto entity.User) error {
	return u.userRepository.CreateUser(entity.User{
		ID:             userDto.ID,
		Username:       userDto.Username,
		Email:          userDto.Email,
		Phone:          userDto.Phone,
		Password:       userDto.Password,
		OpenId:         userDto.OpenId,
		WechatNickname: userDto.WechatNickname,
	})
}

func (u userService) GetUserByUserId(userId string) (entity.User, error) {
	userEntity, err := u.userRepository.GetUserById(userId)
	if err != nil {
		return entity.User{}, err
	}
	return MapEntityUserToUser(userEntity), nil
}

func NewUserService(userRepo repository.UserRepository) UserService {
	return &userService{
		userRepository: userRepo,
	}
}

func MapEntityUserToUser(user entity.User) entity.User {
	return entity.User{
		ID:             user.ID,
		Username:       user.Username,
		Email:          user.Email,
		Phone:          user.Phone,
		Password:       user.Password,
		WechatNickname: user.WechatNickname,
	}
}
