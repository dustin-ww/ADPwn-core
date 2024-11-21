package repository

import (
	"ADPwn/database/project/model"
	"context"

	"github.com/jmoiron/sqlx"
)

type ProjectRepository interface {
	GetAllProjects(ctx context.Context) ([]model.Project, error)
}

type SQLProjectRepository struct {
	DB *sqlx.DB
}

func NewSQLProjectRepository(db *sqlx.DB) *SQLProjectRepository {
	return &SQLProjectRepository{DB: db}
}

func (r *SQLProjectRepository) GetAllProjects(ctx context.Context) ([]model.Project, error) {
	var projects []model.Project
	query := "SELECT * FROM Projects"
	err := r.DB.SelectContext(ctx, &projects, query)
	if err != nil {
		return nil, err
	}
	return projects, nil
}
