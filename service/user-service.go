package service

import "github.com/zhuliminl/easyrn-server/dto"

type UserService interface {
	GetUserById(id string) (dto.User, error)
}

type userService struct {
	// userService
}
