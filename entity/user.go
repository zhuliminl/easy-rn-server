package entity

import "gorm.io/gorm"

type User struct {
	gorm.Model
	ID             string `gorm:"primary_key;type:varchar(64)" json:"user_id"`
	Username       string `gorm:"type:varchar(255)" json:"username"`
	Email          string `gorm:"type:varchar(255)" json:"email"`
	Password       string `gorm:"->;<-;not null" json:"pss"`
	Phone          string `gorm:"type:varchar(255)" json:"phone"`
	OpenId         string `gorm:"type:varchar(255)" json:"wx_openid"`
	WechatNickname string `gorm:"type:varchar(255)" json:"wx_nickname"`
	Projects       []Project
	Teams          []Team
}

/*
func (u *User) BeforeCreate(tx *gorm.DB) (err error) {
	u.ID = uuid.NewV4().String()
	return err
}
*/
