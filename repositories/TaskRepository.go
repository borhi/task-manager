package repositories

import "task-manager/models"

//@TODO add db level
type TaskRepository struct {
}

type ITaskRepository interface {
	Add(task models.TaskModel) (*models.TaskModel, error)
	FindById(id uint) (*models.TaskModel, error)
	FindByColumnId(id uint) ([]models.TaskModel, error)
	Update(task models.TaskModel) (*models.TaskModel, error)
	DeleteById(id uint) error
}

func (repository TaskRepository) Add(task models.TaskModel) (*models.TaskModel, error) {
	task.Id = 1
	return &task, nil
}

func (repository TaskRepository) FindById(id uint) (*models.TaskModel, error) {
	task := &models.TaskModel{
		Id:          id,
		Name:        "test",
		Description: "test",
		Position:    1,
		ColumnId:    1,
	}

	return task, nil
}

func (repository TaskRepository) FindByColumnId(id uint) ([]models.TaskModel, error) {
	tasks := []models.TaskModel{models.TaskModel{
		Id:          1,
		Name:        "test",
		Description: "test",
		Position:    1,
		ColumnId:    id,
	}}

	return tasks, nil
}

func (repository TaskRepository) Update(task models.TaskModel) (*models.TaskModel, error) {
	return &task, nil
}

func (repository TaskRepository) DeleteById(id uint) error {
	return nil
}
