package repository

import (
	"github.com/zhuliminl/easyrn-server/entity"
	"gorm.io/gorm"
)

type UserRepository interface {
	GetUserById(userId string) entity.User
}

type userConnection struct {
	connection *gorm.DB
}

func (u userConnection) GetUserById(userId string) entity.User {
	return entity.User{}
}

func NewUserRepository(db *gorm.DB) UserRepository {
	return &userConnection{
		connection: db,
	}
}
