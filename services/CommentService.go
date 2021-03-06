package services

import (
	"task-manager/models"
	"task-manager/repositories"
)

type CommentService struct {
	Repository repositories.ICommentRepository
}

func (service CommentService) Create(comment models.CommentModel) (*models.CommentModel, error) {
	newComment, err := service.Repository.Add(comment)
	if err != nil {
		return nil, err
	}

	return newComment, nil
}

func (service CommentService) GetById(id int64) (*models.CommentModel, error) {
	return service.Repository.FindById(id)
}

func (service CommentService) GetByTaskId(id int64) ([]*models.CommentModel, error) {
	return service.Repository.FindByTaskId(id)
}

func (service CommentService) Update(comment models.CommentModel) (*models.CommentModel, error) {
	return service.Repository.Update(comment)
}

func (service CommentService) DeleteById(id int64) error {
	return service.Repository.DeleteById(id)
}


