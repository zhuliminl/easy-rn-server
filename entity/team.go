package entity

type Team struct {
	ID     string `gorm:"primary_key:auto_increment" json:"team_id"`
	UserID string `gorm:"not null" json:"creater_id"`
	Title  string `gorm:"type:varchar(255)" json:"title"`
	Desc   string `gorm:"type:text" json:"desc"`
}
