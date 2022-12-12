package repository

import (
	"github.com/zhuliminl/easyrn-server/db"
	"github.com/zhuliminl/easyrn-server/entity"
)

type TeamRepository interface {
	Save(team entity.Team) error
}

type teamRepository struct {
}

func (u teamRepository) Save(team entity.Team) error {
	return db.DB.Create(&team).Error
}

func NewTeamRepository() TeamRepository {
	return &teamRepository{}
}
