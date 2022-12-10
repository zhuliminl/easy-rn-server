package repository

type ProjectRepository interface {
}

type projectRepository struct {
}

func NewProjectRepository() ProjectRepository {
	return &projectRepository{}
}
