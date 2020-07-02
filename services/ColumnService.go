package services

import (
	"task-manager/models"
	"task-manager/repositories"
)

type ColumnService struct {
	Repository repositories.IColumnRepository
}

func (service ColumnService) Create(column models.ColumnModel) (*models.ColumnModel, error) {
	newColumn, err := service.Repository.Add(column)
	if err != nil {
		return nil, err
	}

	return newColumn, nil
}

func (service ColumnService) GetById(id int64) (*models.ColumnModel, error) {
	return service.Repository.FindById(id)
}

func (service ColumnService) GetByProjectId(id int64) ([]*models.ColumnModel, error) {
	return service.Repository.FindByProjectId(id)
}

func (service ColumnService) Update(column models.ColumnModel) (*models.ColumnModel, error) {
	return service.Repository.Update(column)
}

func (service ColumnService) DeleteById(id int64) error {
	return service.Repository.DeleteById(id)
}
