package repositories

import "task-manager/models"

//@TODO add db level
type ProjectRepository struct {
}

type IProjectRepository interface {
	Add(project models.ProjectModel) (*models.ProjectModel, error)
	FindById(id uint) (*models.ProjectModel, error)
	FindAll() ([]models.ProjectModel, error)
	Update(project models.ProjectModel) (*models.ProjectModel, error)
	DeleteById(id uint) error
}

func (repository ProjectRepository) Add(project models.ProjectModel) (*models.ProjectModel, error) {
	project.Id = 1
	return &project, nil
}

func (repository ProjectRepository) FindById(id uint) (*models.ProjectModel, error) {
	project := &models.ProjectModel{
		Id:          1,
		Name:        "test",
		Description: "test",
	}

	return project, nil
}

func (repository ProjectRepository) FindAll() ([]models.ProjectModel, error) {
	projects := []models.ProjectModel{models.ProjectModel{
		Id:          1,
		Name:        "test",
		Description: "test",
	}, models.ProjectModel{
		Id:          2,
		Name:        "test2",
		Description: "test2",
	}}

	return projects, nil
}

func (repository ProjectRepository) Update(project models.ProjectModel) (*models.ProjectModel, error) {
	return &project, nil
}

func (repository ProjectRepository) DeleteById(id uint) error {
	return nil
}
