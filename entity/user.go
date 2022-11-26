package entity

import (
	uuid "github.com/satori/go.uuid"
	"gorm.io/gorm"
)

type User struct {
	ID   string `json:"id" gorm:"type:varchar(64)"`
	Name string `gorm:"type:varchar(255)" json:"name"`
	//Email    string `gorm:"uniqueIndex;type:varchar(255)" json:"email"`
	Email    string `gorm:"type:varchar(255)" json:"email"`
	Password string `gorm:"->;<-;not null" json:"-"`
	Phone    string `gorm:"type:varchar(255)" json:"phone"`
}

func (u *User) BeforeCreate(tx *gorm.DB) (err error) {
	u.ID = uuid.NewV4().String()
	return err
}
