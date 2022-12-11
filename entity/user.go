package entity

import "gorm.io/gorm"

type User struct {
	gorm.Model
	ID             string `gorm:"primary_key;type:varchar(64)"`
	Username       string `gorm:"type:varchar(255)"`
	Email          string `gorm:"type:varchar(255)"`
	Password       string `gorm:"->;<-;not null"`
	Phone          string `gorm:"type:varchar(255)"`
	OpenId         string `gorm:"type:varchar(255)"`
	WechatNickname string `gorm:"type:varchar(255)"`
	Projects       []Project
	Teams          []Team
}

/*
func (u *User) BeforeCreate(tx *gorm.DB) (err error) {
	u.ID = uuid.NewV4().String()
	return err
}
*/
