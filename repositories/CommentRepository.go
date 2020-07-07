package repositories

import (
	"errors"
	"fmt"
	"task-manager/adapters"
	"task-manager/models"
)

//@TODO add db level
type CommentRepository struct {
	adapters.IDbAdapter
}

type ICommentRepository interface {
	Add(comment models.CommentModel) (*models.CommentModel, error)
	FindById(id int64) (*models.CommentModel, error)
	FindByTaskId(id int64) ([]*models.CommentModel, error)
	Update(comment models.CommentModel) (*models.CommentModel, error)
	DeleteById(id int64) error
}

func (repository CommentRepository) Add(comment models.CommentModel) (*models.CommentModel, error) {
	rows, err := repository.Query(fmt.Sprintf(
		"INSERT INTO comment (text, created_at, task_id) values ('%s', current_timestamp, %d) " +
			"RETURNING id, text, created_at, task_id",
		comment.Text, comment.TaskId,
	))
	if err != nil {
		return nil, err
	}

	newComment := new(models.CommentModel)
	rows.Next()
	err = rows.Scan(&newComment.Id, &newComment.Text, &newComment.CreatedAt, &newComment.TaskId)
	if err != nil {
		return nil, err
	}
	return newComment, nil
}

func (repository CommentRepository) FindById(id int64) (*models.CommentModel, error) {
	rows, err := repository.Query(fmt.Sprintf("SELECT * FROM comment WHERE id = %d", id))
	if err != nil {
		return nil, err
	}

	comment := new(models.CommentModel)
	rows.Next()
	if err := rows.Scan(&comment.Id, &comment.Text, &comment.CreatedAt, &comment.TaskId); err != nil {
		return nil, errors.New("not found")
	}

	return comment, nil
}

func (repository CommentRepository) FindByTaskId(id int64) ([]*models.CommentModel, error) {
	rows, err := repository.Query(fmt.Sprintf(
		"SELECT * FROM comment WHERE task_id = %d ORDER BY created_at DESC",
		id,
	))
	if err != nil {
		return nil, err
	}

	comments := make([]*models.CommentModel, 0)
	for rows.Next() {
		comment := new(models.CommentModel)
		err := rows.Scan(&comment.Id, &comment.Text, &comment.CreatedAt, &comment.TaskId)
		if err != nil {
			return nil, err
		}
		comments = append(comments, comment)
	}

	return comments, nil
}

func (repository CommentRepository) Update(comment models.CommentModel) (*models.CommentModel, error) {
	rows, err := repository.Query(fmt.Sprintf(
		"UPDATE comment SET text='%s', task_id=%d WHERE id=%d RETURNING id, text, created_at, task_id",
		comment.Text, comment.TaskId, comment.Id,
	))
	if err != nil {
		return nil, err
	}

	updatedComment := new(models.CommentModel)
	rows.Next()
	err = rows.Scan(&updatedComment.Id, &updatedComment.Text, &updatedComment.CreatedAt, &updatedComment.TaskId)
	if err != nil {
		return nil, errors.New("not found")
	}

	return updatedComment, nil
}

func (repository CommentRepository) DeleteById(id int64) error {
	res, err := repository.Execute(fmt.Sprintf("DELETE FROM comment WHERE id=%d", id))
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
