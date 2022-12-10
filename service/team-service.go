package service

import (
	"github.com/zhuliminl/easyrn-server/repository"
)

type TeamService interface {
}

type teamService struct {
	userRepository repository.UserRepository
}

func NewTeamService(userRepo repository.UserRepository) TeamService {
	return &userService{
		userRepository: userRepo,
	}
}
