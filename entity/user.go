package entity

import "gorm.io/gorm"

type User struct {
	gorm.Model
	ID             string `json:"id" gorm:"type:varchar(64)"`
	Username       string `gorm:"type:varchar(255)" json:"username"`
	Email          string `gorm:"type:varchar(255)" json:"email"`
	Password       string `gorm:"->;<-;not null" json:"-"`
	Phone          string `gorm:"type:varchar(255)" json:"phone"`
	OpenId         string `gorm:"type:varchar(255)" json:"openId"`
	WechatNickname string `gorm:"type:varchar(255)" json:"wechat_nickname"`
	Projects       []Project
}

/*
func (u *User) BeforeCreate(tx *gorm.DB) (err error) {
	u.ID = uuid.NewV4().String()
	return err
}
*/
