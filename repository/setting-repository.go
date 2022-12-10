package repository

type SettingRepository interface {
}

type settingRepository struct {
}

func NewSettingRepository() SettingRepository {
	return &settingRepository{}
}
