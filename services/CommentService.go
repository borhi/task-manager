package services

import (
	"task-manager/models"
	"task-manager/repositories"
)

type CommentService struct {
	repository repositories.ICommentRepository
}

func (service CommentService) Create(comment models.CommentModel) (*models.CommentModel, error) {
	newComment, err := service.repository.Create(comment)
	if err != nil {
		return nil, err
	}

	return newComment, nil
}

func (service CommentService) GetById(id uint) (*models.CommentModel, error) {
	return service.repository.GetById(id)
}

func (service CommentService) GetByTaskId(id uint) ([]models.CommentModel, error) {
	return service.repository.GetByTaskId(id)
}

func (service CommentService) Update(comment models.CommentModel) (*models.CommentModel, error) {
	return service.repository.Update(comment)
}

func (service CommentService) DeleteById(id uint) error {
	return service.repository.DeleteById(id)
}


