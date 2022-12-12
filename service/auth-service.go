package service

import (
	uuid "github.com/satori/go.uuid"
	"github.com/zhuliminl/easyrn-server/constError"
	"github.com/zhuliminl/easyrn-server/dto"
	"github.com/zhuliminl/easyrn-server/entity"
	"github.com/zhuliminl/easyrn-server/helper"
	"github.com/zhuliminl/easyrn-server/repository"
)

type AuthService interface {
	VerifyCredentialByEmail(email string, password string) (entity.User, error)
	VerifyCredentialByPhone(phone string, password string) (entity.User, error)

	VerifyRegisterByEmail(user dto.UserRegisterByEmail) error
	VerifyRegisterByPhone(user dto.UserRegisterByPhone) error

	CreateUserByEmail(user dto.UserRegisterByEmail) (entity.User, error)
	CreateUserByPhone(user dto.UserRegisterByPhone) (entity.User, error)
}

type authService struct {
	userRepository repository.UserRepository
	userService    UserService
	jwtService     JWTService
}

func (a authService) VerifyCredentialByEmail(email string, password string) (entity.User, error) {
	var user entity.User
	// fixme
	return user, nil
}

func (a authService) VerifyCredentialByPhone(phone string, password string) (entity.User, error) {
	var user entity.User
	// fixme
	return user, nil
}

func (a authService) CreateUserByEmail(userRegister dto.UserRegisterByEmail) (entity.User, error) {
	username := userRegister.Username
	if username == "" {
		username = helper.GenerateDefaultUserName()
	}
	var user = entity.User{
		ID:       uuid.NewV4().String(),
		Username: username,
		Email:    userRegister.Email,
		Password: userRegister.Password,
	}
	err := a.userService.CreateUser(user)
	if err != nil {
		return entity.User{}, err
	}
	return user, nil
}

func (a authService) CreateUserByPhone(userRegister dto.UserRegisterByPhone) (entity.User, error) {
	username := userRegister.Username
	if username == "" {
		username = helper.GenerateDefaultUserName()
	}
	var user = entity.User{
		ID:       uuid.NewV4().String(),
		Username: username,
		Phone:    userRegister.Phone,
		Password: userRegister.Password,
	}
	err := a.userService.CreateUser(user)
	if err != nil {
		return entity.User{}, err
	}
	return user, nil
}

func (a authService) VerifyRegisterByEmail(user dto.UserRegisterByEmail) error {
	if !helper.IsEmailValid(user.Email) {
		return constError.NewEmailNotValid(nil, "邮箱格式错误")
	}
	if !helper.IsPasswordValid(user.Password) {
		return constError.NewPasswordNotValid(nil, "密码格式错误")
	}

	_, err := a.userRepository.GetUserByEmail(user.Email)
	if constError.Is(err, constError.UserNotFound) {
		return nil
	}
	if err != nil {
		return err
	}
	return constError.NewUserDuplicated(nil, "用户已注册")
}

func (a authService) VerifyRegisterByPhone(user dto.UserRegisterByPhone) error {
	if !helper.IsPhoneValid(user.Phone) {
		return constError.NewPhoneNumberNotValid(nil, "手机格式错误")
	}
	if !helper.IsPasswordValid(user.Password) {
		return constError.NewPasswordNotValid(nil, "密码格式错误")
	}

	_, err := a.userRepository.GetUserByPhone(user.Phone)
	if constError.Is(err, constError.UserNotFound) {
		return nil
	}
	if err != nil {
		return err
	}
	return constError.NewUserDuplicated(nil, "用户已注册")
}

func NewAuthService(userRepo repository.UserRepository, userService UserService, jwtService JWTService) AuthService {
	return &authService{
		userRepository: userRepo,
		userService:    userService,
		jwtService:     jwtService,
	}
}
