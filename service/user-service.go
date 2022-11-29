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
	return u.userRepository.CreateUser(entity.User{
		ID:             userDto.UserId,
		Username:       userDto.Username,
		Email:          userDto.Email,
		Phone:          userDto.Phone,
		Password:       userDto.Password,
		OpenId:         userDto.OpenId,
		WechatNickname: userDto.WechatNickname,
	})
}

func (u userService) GetUserByUserId(userId string) (dto.User, error) {
	userEntity, err := u.userRepository.GetUserById(userId)
	if err != nil {
		return dto.User{}, err
	}
	return MapEntityUserToUser(userEntity), nil
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
		Password:       user.Password,
		WechatNickname: user.WechatNickname,
	}
}
