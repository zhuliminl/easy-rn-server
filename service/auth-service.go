package service

import (
	"github.com/zhuliminl/easyrn-server/constError"
	"github.com/zhuliminl/easyrn-server/dto"
	"github.com/zhuliminl/easyrn-server/helper"
	"github.com/zhuliminl/easyrn-server/repository"
)

type AuthService interface {
	VerifyRegisterByEmail(user dto.UserRegisterByEmail) error
	VerifyRegisterByPhone(user dto.UserRegisterByPhone) error
}

type authService struct {
	userRepository repository.UserRepository
}

func (a authService) VerifyRegisterByEmail(user dto.UserRegisterByEmail) error {
	if !helper.IsEmailValid(user.Email) {
		return constError.NewEmailNotValid(nil, "邮箱格式错误")
	}
	return nil
}

func (a authService) VerifyRegisterByPhone(user dto.UserRegisterByPhone) error {
	return nil
}

func NewAuthService(userRepo repository.UserRepository) AuthService {
	return &authService{
		userRepository: userRepo,
	}
}
