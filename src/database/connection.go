package main

import (
	"ADPwn/database/project/repository"
	"ADPwn/database/project/service"
	"fmt"
	"log"

	"github.com/jmoiron/sqlx"
	_ "github.com/mattn/go-sqlite3"
)

func main() {
	// Datenbankverbindung initialisieren
	db, err := sqlx.Connect("sqlite3", "./db.sqlite")
	if err != nil {
		log.Fatalf("Fehler beim Herstellen der Verbindung zur Datenbank: %v", err)
	}
	defer db.Close()

	// Repository und Service initialisieren
	projectRepo := repository.NewSQLProjectRepository(db)
	projectService := service.NewProjectService(projectRepo)

	// Abrufen und Ausgeben der Projekte
	projects, err := projectService.GetAllProjects()
	if err != nil {
		fmt.Printf("Fehler beim Abrufen der Projekte: %v\n", err)
		return
	}

	// Projekte ausgeben
	for _, project := range projects {
		fmt.Printf("Projekt: %s (UUID: %s)\n", project.Name, project.UUID)
	}
}
