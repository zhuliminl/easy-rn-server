package repository

type TeamRepository interface {
}

type teamRepository struct {
}

func NewTeamRepository() TeamRepository {
	return &teamRepository{}
}
