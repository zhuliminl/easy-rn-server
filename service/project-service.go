package service

import (
	"github.com/zhuliminl/easyrn-server/repository"
)

type ProjectService interface {
}

type projectService struct {
	userRepository repository.UserRepository
}

func NewProjectService(userRepo repository.UserRepository) ProjectService {
	return &userService{
		userRepository: userRepo,
	}
}
