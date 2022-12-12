package service

import (
	uuid "github.com/satori/go.uuid"
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
	team := entity.Team{
		ID:     uuid.NewV4().String(),
		UserID: useId,
		Title:  teamCreate.Title,
		Desc:   teamCreate.Desc,
	}
	return t.teamRepository.Save(team)
}

func NewTeamService(userRepo repository.UserRepository, teamRepository repository.TeamRepository) TeamService {
	return &teamService{
		userRepository: userRepo,
		teamRepository: teamRepository,
	}
}
