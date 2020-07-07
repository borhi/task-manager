package repositories

import (
	"errors"
	"fmt"
	"task-manager/adapters"
	"task-manager/models"
)

//@TODO add db level
type TaskRepository struct {
	adapters.IDbAdapter
}

type ITaskRepository interface {
	Add(task models.TaskModel) (*models.TaskModel, error)
	FindById(id int64) (*models.TaskModel, error)
	FindByColumnId(id int64) ([]*models.TaskModel, error)
	Update(task models.TaskModel) (*models.TaskModel, error)
	DeleteById(id int64) error
}

func (repository TaskRepository) Add(task models.TaskModel) (*models.TaskModel, error) {
	rows, err := repository.Query(fmt.Sprintf(
		"INSERT INTO task (name, description, position, column_id) values ('%s', '%s', %d, %d) "+
			"RETURNING id, name, description, position, column_id",
		task.Name, task.Description, task.Position, task.ColumnId,
	))
	if err != nil {
		return nil, err
	}

	newTask := new(models.TaskModel)
	rows.Next()
	err = rows.Scan(
		&newTask.Id,
		&newTask.Name,
		&newTask.Description,
		&newTask.Position,
		&newTask.ColumnId,
	)
	if err != nil {
		return nil, err
	}

	return newTask, nil
}

func (repository TaskRepository) FindById(id int64) (*models.TaskModel, error) {
	rows, err := repository.Query(fmt.Sprintf("SELECT * FROM task WHERE id = %d", id))
	if err != nil {
		return nil, err
	}

	task := new(models.TaskModel)
	rows.Next()
	if err := rows.Scan(&task.Id, &task.Name, &task.Description, &task.Position, &task.ColumnId); err != nil {
		return nil, errors.New("not found")
	}

	return task, nil
}

func (repository TaskRepository) FindByColumnId(id int64) ([]*models.TaskModel, error) {
	rows, err := repository.Query(fmt.Sprintf("SELECT * FROM task WHERE column_id = %d ORDER BY position", id))
	if err != nil {
		return nil, err
	}

	tasks := make([]*models.TaskModel, 0)
	for rows.Next() {
		task := new(models.TaskModel)
		err := rows.Scan(&task.Id, &task.Name, &task.Description, &task.Position, &task.ColumnId)
		if err != nil {
			return nil, err
		}
		tasks = append(tasks, task)
	}

	return tasks, nil
}

func (repository TaskRepository) Update(task models.TaskModel) (*models.TaskModel, error) {
	rows, err := repository.Query(fmt.Sprintf(
		"UPDATE task SET name='%s', description='%s', position=%d, column_id=%d WHERE id=%d "+
			"RETURNING id, name, description, position, column_id",
		task.Name, task.Description, task.Position, task.ColumnId, task.Id,
	))
	if err != nil {
		return nil, err
	}

	updatedTask := new(models.TaskModel)
	rows.Next()
	err = rows.Scan(
		&updatedTask.Id,
		&updatedTask.Name,
		&updatedTask.Description,
		&updatedTask.Position,
		&updatedTask.ColumnId,
	)
	if err != nil {
		return nil, errors.New("not found")
	}

	return updatedTask, nil
}

func (repository TaskRepository) DeleteById(id int64) error {
	res, err := repository.Execute(fmt.Sprintf("DELETE FROM task WHERE id=%d", id))
	if err != nil {
		return err
	}

	count, err := res.RowsAffected()
	if err != nil {
		return err
	}

	if count == 0 {
		return errors.New("not found")
	}

	return nil
}
