package repositories

import (
	"errors"
	"fmt"
	"task-manager/adapters"
	"task-manager/models"
)

type ColumnRepository struct {
	adapters.IDbAdapter
}

type IColumnRepository interface {
	Add(column models.ColumnModel) (*models.ColumnModel, error)
	FindById(id int64) (*models.ColumnModel, error)
	FindByProjectId(id int64) ([]*models.ColumnModel, error)
	Update(column models.ColumnModel) (*models.ColumnModel, error)
	DeleteById(id int64) error
}

func (repository ColumnRepository) Add(column models.ColumnModel) (*models.ColumnModel, error) {
	rows, err := repository.Query(fmt.Sprintf(
		"INSERT INTO \"column\" (name, position, project_id) values ('%s', %d, %d) RETURNING id, name, position, project_id",
		column.Name, column.Position, column.ProjectId,
	))
	if err != nil {
		return nil, err
	}

	rows.Next()
	newColumn := new(models.ColumnModel)
	err = rows.Scan(&newColumn.Id, &newColumn.Name, &newColumn.Position, &newColumn.ProjectId)
	if err != nil {
		return nil, err
	}
	return newColumn, nil
}

func (repository ColumnRepository) FindById(id int64) (*models.ColumnModel, error) {
	rows, err := repository.Query(fmt.Sprintf("SELECT * FROM \"column\" WHERE id = %d", id))
	if err != nil {
		return nil, err
	}

	column := new(models.ColumnModel)
	rows.Next()
	if err := rows.Scan(&column.Id, &column.Name, &column.Position, &column.ProjectId); err != nil {
		return nil, errors.New("not found")
	}

	return column, nil
}

func (repository ColumnRepository) FindByProjectId(id int64) ([]*models.ColumnModel, error) {
	rows, err := repository.Query(fmt.Sprintf("SELECT * FROM \"column\" WHERE project_id = %d", id))
	if err != nil {
		return nil, err
	}

	columns := make([]*models.ColumnModel, 0)
	for rows.Next() {
		column := new(models.ColumnModel)
		err := rows.Scan(&column.Id, &column.Name, &column.Position, &column.ProjectId)
		if err != nil {
			return nil, err
		}
		columns = append(columns, column)
	}

	return columns, nil
}

func (repository ColumnRepository) Update(column models.ColumnModel) (*models.ColumnModel, error) {
	rows, err := repository.Query(fmt.Sprintf(
		"UPDATE \"column\" SET name='%s', position=%d, project_id=%d WHERE id=%d " +
			"RETURNING id, name, position, project_id;",
		column.Name, column.Position, column.ProjectId, column.Id,
	))
	if err != nil {
		return nil, err
	}

	updatedColumn := new(models.ColumnModel)
	rows.Next()
	err = rows.Scan(&updatedColumn.Id, &updatedColumn.Name, &updatedColumn.Position, &updatedColumn.ProjectId)
	if err != nil {
		return nil, errors.New("not found")
	}

	return updatedColumn, nil
}

func (repository ColumnRepository) DeleteById(id int64) error {
	res, err := repository.Execute(fmt.Sprintf("DELETE FROM \"column\" WHERE id=%d", id))
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
