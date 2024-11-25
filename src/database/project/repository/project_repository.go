package repository

import (
	"ADPwn/database/project/model"
	"context"

	"github.com/jmoiron/sqlx"
)

type ProjectRepository interface {
	AllProjects(ctx context.Context) ([]model.Project, error)
	SaveProject(ctx context.Context, project model.Project) error
}

type SQLProjectRepository struct {
	DB *sqlx.DB
}

func NewSQLProjectRepository(db *sqlx.DB) *SQLProjectRepository {
	return &SQLProjectRepository{DB: db}
}

func (r *SQLProjectRepository) AllProjects(ctx context.Context) ([]model.Project, error) {
	var projects []model.Project
	query := "SELECT * FROM Projects"
	err := r.DB.SelectContext(ctx, &projects, query)
	if err != nil {
		return nil, err
	}
	return projects, nil
}

func (r *SQLProjectRepository) SaveProject(ctx context.Context, project model.Project) error {
	_, err := r.DB.NamedExecContext(ctx, `INSERT INTO project (first_name, last_name, email)
        VALUES (:first_name, :last_name, :email)`, project)
	if err != nil {
		return err
	}
	return nil
}
