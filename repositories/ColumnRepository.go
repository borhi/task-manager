package repositories

import "task-manager/models"

//@TODO add db level
type ColumnRepository struct {
}

type IColumnRepository interface {
	Add(column models.ColumnModel) (*models.ColumnModel, error)
	FindById(id uint) (*models.ColumnModel, error)
	FindByProjectId(id uint) ([]models.ColumnModel, error)
	Update(column models.ColumnModel) (*models.ColumnModel, error)
	DeleteById(id uint) error
}

func (repository ColumnRepository) Add(column models.ColumnModel) (*models.ColumnModel, error) {
	column.Id = 1
	return &column, nil
}

func (repository ColumnRepository) FindById(id uint) (*models.ColumnModel, error) {
	column := &models.ColumnModel{
		Id:        id,
		Name:      "test",
		Position:  1,
		ProjectId: 1,
	}

	return column, nil
}

func (repository ColumnRepository) FindByProjectId(id uint) ([]models.ColumnModel, error) {
	columns := []models.ColumnModel{models.ColumnModel{
		Id:        1,
		Name:      "test",
		Position:  1,
		ProjectId: id,
	}}

	return columns, nil
}

func (repository ColumnRepository) Update(column models.ColumnModel) (*models.ColumnModel, error) {
	return &column, nil
}

func (repository ColumnRepository) DeleteById(id uint) error {
	return nil
}
