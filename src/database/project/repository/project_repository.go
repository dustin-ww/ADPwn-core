package repository

import (
	"ADPwn/database/project/model"

	"github.com/jmoiron/sqlx"
)

// ProjectRepository definiert die Methoden für den Datenbankzugriff.
type ProjectRepository interface {
	GetAllProjects() ([]model.Project, error)
}

// SQLProjectRepository ist die konkrete Implementierung von ProjectRepository für SQL-Datenbanken.
type SQLProjectRepository struct {
	DB *sqlx.DB
}

// NewSQLProjectRepository erstellt eine neue Instanz des SQL-basierten Repositories.
func NewSQLProjectRepository(db *sqlx.DB) *SQLProjectRepository {
	return &SQLProjectRepository{DB: db}
}

// GetAllProjects liest alle Projekte aus der Datenbank.
func (r *SQLProjectRepository) GetAllProjects() ([]model.Project, error) {
	var projects []model.Project
	err := r.DB.Select(&projects, "SELECT * FROM projects")
	if err != nil {
		return nil, err
	}
	return projects, nil
}
