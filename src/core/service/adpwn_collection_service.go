package service

import (
	"ADPwn/core/internal/db"
	"ADPwn/core/internal/repository"
	"ADPwn/core/model/adpwn"
	"context"
	"fmt"
	"gorm.io/gorm"
	"log"
)

type ADPwnCollectionService struct {
	db                  *gorm.DB
	adpwnModuleRepo     repository.ADPwnModuleRepository
	adpwnCollectionRepo repository.ADPwnCollectionRepository
}

func NewADPwnCollectionService() (*ADPwnCollectionService, error) {
	db, err := db.GetPostgresDB()
	if err != nil {
		return nil, fmt.Errorf("failed to get database connection: %w", err)
	}

	return &ADPwnCollectionService{
		db:                  db,
		adpwnModuleRepo:     repository.NewPostgresADPwnModuleRepository(),
		adpwnCollectionRepo: repository.NewPostgresADPwnCollectionRepository(),
	}, nil
}

func (s *ADPwnCollectionService) GetAllCollectionsWithModules(ctx context.Context) ([]*adpwn.Collection, error) {
	return db.ExecutePostgresRead(ctx, s.db, func(db *gorm.DB) ([]*adpwn.Collection, error) {
		collections, err := s.adpwnCollectionRepo.GetAll(ctx, db)
		if err != nil {
			return nil, err
		}

		for _, collection := range collections {
			modules, err := s.adpwnCollectionRepo.GetModulesForCollection(ctx, db, collection.ID)
			if err != nil {
				return nil, err
			}

			collection.Modules = modules
		}

		if len(collections) == 0 {
			log.Println("no collections found in database")
		}
		return collections, nil
	})
}
