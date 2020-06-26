package services

import (
	"task-manager/models"
	"task-manager/repositories"
)

type ProjectService struct {
	ProjectRepository repositories.IProjectRepository
	ColumnRepository  repositories.IColumnRepository
}

func (service ProjectService) Create(project models.ProjectModel) (*models.ProjectModel, error) {
	newProject, err := service.ProjectRepository.Create(project)
	if err != nil {
		return nil, err
	}

	column := models.ColumnModel{
		Name: "default",
		Position: 1,
		ProjectId: newProject.Id,
	}

	_, err = service.ColumnRepository.Create(column)
	if err != nil {
		return nil, err
	}

	return newProject, nil
}

func (service ProjectService) GetById(id uint) (*models.ProjectModel, error) {
	return service.ProjectRepository.GetById(id)
}

func (service ProjectService) GetList() ([]models.ProjectModel, error) {
	return service.ProjectRepository.GetList()
}

func (service ProjectService) Update(project models.ProjectModel) (*models.ProjectModel, error) {
	return service.ProjectRepository.Update(project)
}

func (service ProjectService) DeleteById(id uint) error {
	return service.ProjectRepository.DeleteById(id)
}
