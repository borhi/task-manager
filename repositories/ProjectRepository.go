package repositories

import (
	"errors"
	"fmt"
	"task-manager/adapters"
	"task-manager/models"
)

type ProjectRepository struct {
	adapters.IDbAdapter
}

type IProjectRepository interface {
	Add(project models.ProjectModel) (*models.ProjectModel, error)
	FindById(id int64) (*models.ProjectModel, error)
	FindAll() ([]*models.ProjectModel, error)
	Update(project models.ProjectModel) (*models.ProjectModel, error)
	DeleteById(id int64) error
}

func (repository ProjectRepository) Add(project models.ProjectModel) (*models.ProjectModel, error) {
	rows, err := repository.Query(fmt.Sprintf(
		"INSERT INTO project (name, description) values ('%s', '%s') RETURNING id, name, description",
		project.Name, project.Description,
	))
	if err != nil {
		return nil, err
	}

	rows.Next()
	newProject := new(models.ProjectModel)
	if err = rows.Scan(&newProject.Id, &newProject.Name, &newProject.Description); err != nil {
		return nil, err
	}
	return newProject, nil
}

func (repository ProjectRepository) FindById(id int64) (*models.ProjectModel, error) {
	rows, err := repository.Query(fmt.Sprintf("SELECT * FROM project WHERE id = %d", id))
	if err != nil {
		return nil, err
	}

	project := new(models.ProjectModel)
	rows.Next()
	if err := rows.Scan(&project.Id, &project.Name, &project.Description); err != nil {
		return nil, errors.New("not found")
	}

	return project, nil
}

func (repository ProjectRepository) FindAll() ([]*models.ProjectModel, error) {
	rows, err := repository.Query("SELECT * FROM project")
	if err != nil {
		return nil, err
	}

	projects := make([]*models.ProjectModel, 0)
	for rows.Next() {
		project := new(models.ProjectModel)
		err := rows.Scan(&project.Id, &project.Name, &project.Description)
		if err != nil {
			return nil, err
		}
		projects = append(projects, project)
	}

	return projects, nil
}

func (repository ProjectRepository) Update(project models.ProjectModel) (*models.ProjectModel, error) {
	rows, err := repository.Query(fmt.Sprintf(
		"UPDATE project SET name='%s', description='%s' WHERE id=%d RETURNING id, name, description;",
		project.Name, project.Description, project.Id,
	))
	if err != nil {
		return nil, err
	}

	updatedProject := new(models.ProjectModel)
	rows.Next()
	if err := rows.Scan(&updatedProject.Id, &updatedProject.Name, &updatedProject.Description); err != nil {
		return nil, errors.New("not found")
	}

	return updatedProject, nil
}

func (repository ProjectRepository) DeleteById(id int64) error {
	res, err := repository.Execute(fmt.Sprintf("DELETE FROM project WHERE id=%d", id))
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
