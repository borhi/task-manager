package services

import (
	"task-manager/models"
	"task-manager/repositories"
)

type TaskService struct {
	Repository repositories.ITaskRepository
}

func (service TaskService) Create(task models.TaskModel) (*models.TaskModel, error) {
	newTask, err := service.Repository.Add(task)
	if err != nil {
		return nil, err
	}

	return newTask, nil
}

func (service TaskService) GetById(id uint) (*models.TaskModel, error) {
	return service.Repository.FindById(id)
}

func (service TaskService) GetByColumnId(id uint) ([]models.TaskModel, error) {
	return service.Repository.FindByColumnId(id)
}

func (service TaskService) Update(task models.TaskModel) (*models.TaskModel, error) {
	return service.Repository.Update(task)
}

func (service TaskService) DeleteById(id uint) error {
	return service.Repository.DeleteById(id)
}
