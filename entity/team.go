package entity

type Team struct {
	ID     uint64 `gorm:"primary_key:auto_increment"`
	Title  string `gorm:"type:varchar(255)"`
	Desc   string `gorm:"type:text"`
	UserID uint64 `gorm:"not null"`
	User   User   `gorm:"foreignkey:UserID;constraint:onUpdate:CASCADE,onDelete:CASCADE"`
}
