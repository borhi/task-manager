package services

import (
	"task-manager/models"
	"task-manager/repositories"
)

type TaskService struct {
	repository repositories.ITaskRepository
}

func (service TaskService) Create(task models.TaskModel) (*models.TaskModel, error) {
	newTask, err := service.repository.Create(task)
	if err != nil {
		return nil, err
	}

	return newTask, nil
}

func (service TaskService) GetById(id uint) (*models.TaskModel, error) {
	return service.repository.GetById(id)
}

func (service TaskService) GetByColumnId(id uint) ([]models.TaskModel, error) {
	return service.repository.GetByColumnId(id)
}

func (service TaskService) Update(task models.TaskModel) (*models.TaskModel, error) {
	return service.repository.Update(task)
}

func (service TaskService) DeleteById(id uint) error {
	return service.repository.DeleteById(id)
}
