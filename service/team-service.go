package service

import (
	"github.com/zhuliminl/easyrn-server/dto"
	"github.com/zhuliminl/easyrn-server/entity"
	"github.com/zhuliminl/easyrn-server/repository"
)

type TeamService interface {
	CreateTeam(useId string, teamCreate dto.TeamCreate) error
}

type teamService struct {
	userRepository repository.UserRepository
	teamRepository repository.TeamRepository
}

func (t teamService) CreateTeam(useId string, teamCreate dto.TeamCreate) error {
	return t.teamRepository.CreateTeam(entity.Team{})
}

func NewTeamService(userRepo repository.UserRepository, teamRepository repository.TeamRepository) TeamService {
	return &teamService{
		userRepository: userRepo,
		teamRepository: teamRepository,
	}
}
