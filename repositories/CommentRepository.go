package repositories

import (
	"task-manager/models"
	"time"
)

//@TODO add db level
type CommentRepository struct {
}

type ICommentRepository interface {
	Create(comment models.CommentModel) (*models.CommentModel, error)
	GetById(id uint) (*models.CommentModel, error)
	GetByTaskId(id uint) ([]models.CommentModel, error)
	Update(comment models.CommentModel) (*models.CommentModel, error)
	DeleteById(id uint) error
}

func (repository CommentRepository) Create(comment models.CommentModel) (*models.CommentModel, error) {
	comment.Id = 1
	return &comment, nil
}

func (repository CommentRepository) GetById(id uint) (*models.CommentModel, error) {
	task := &models.CommentModel{
		Id:        id,
		Text:      "test",
		CreatedAt: time.Now(),
		TaskId:    1,
	}

	return task, nil
}

func (repository CommentRepository) GetByTaskId(id uint) ([]models.CommentModel, error) {
	comments := []models.CommentModel{models.CommentModel{
		Id:        1,
		Text:      "test",
		CreatedAt: time.Now(),
		TaskId:    id,
	}}

	return comments, nil
}

func (repository CommentRepository) Update(comment models.CommentModel) (*models.CommentModel, error) {
	return &comment, nil
}

func (repository CommentRepository) DeleteById(id uint) error {
	return nil
}
