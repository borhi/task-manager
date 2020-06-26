package services

import (
	"task-manager/models"
	"task-manager/repositories"
)

type ColumnService struct {
	repository repositories.IColumnRepository
}

func (service ColumnService) Create(column models.ColumnModel) (*models.ColumnModel, error) {
	newColumn, err := service.repository.Create(column)
	if err != nil {
		return nil, err
	}

	return newColumn, nil
}

func (service ColumnService) GetById(id uint) (*models.ColumnModel, error) {
	return service.repository.GetById(id)
}

func (service ColumnService) GetByProjectId(id uint) ([]models.ColumnModel, error) {
	return service.repository.GetByProjectId(id)
}

func (service ColumnService) Update(column models.ColumnModel) (*models.ColumnModel, error) {
	return service.repository.Update(column)
}

func (service ColumnService) DeleteById(id uint) error {
	return service.repository.DeleteById(id)
}
