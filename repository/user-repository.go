package repository

import (
	"github.com/zhuliminl/easyrn-server/db"
	"github.com/zhuliminl/easyrn-server/entity"
)

type UserRepository interface {
	GetUserById(userId string) (entity.User, error)
	GetUserByEmail(email string) (entity.User, error)
	CreateUser(user entity.User) error
}

type userRepository struct {
}

func (u userRepository) GetUserByEmail(email string) (entity.User, error) {
	var user entity.User
	if err := db.DB.Where("email = ?", email).First(&user).Error; err != nil {
		return user, err
	}
	return user, nil
}

func (u userRepository) CreateUser(user entity.User) error {
	return db.DB.Create(&user).Error
}

func (u userRepository) GetUserById(userId string) (entity.User, error) {
	var user entity.User
	if err := db.DB.Where("user_id = ?", userId).First(&user).Error; err != nil {
		return user, err
	}
	return user, nil
}

func NewUserRepository() UserRepository {
	return &userRepository{}
}
