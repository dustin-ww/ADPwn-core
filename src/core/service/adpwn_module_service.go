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

type ADPwnModuleService struct {
	db              *gorm.DB
	adpwnModuleRepo repository.ADPwnModuleRepository
}

func NewADPwnModuleService() (*ADPwnModuleService, error) {
	db, err := db.GetPostgresDB()
	if err != nil {
		return nil, fmt.Errorf("failed to get database connection: %w", err)
	}

	return &ADPwnModuleService{
		db:              db,
		adpwnModuleRepo: repository.NewPostgresADPwnModuleRepository(),
	}, nil
}

func (s *ADPwnModuleService) CreateWithObject(ctx context.Context, module *adpwn.Module) (string, error) {
	var attackID string
	err := db.ExecutePostgresInTransaction(ctx, s.db, func(tx *gorm.DB) error {

		// Check if module already exists
		exists, err := s.adpwnModuleRepo.CheckIfExistsByKey(ctx, tx, module.Key)
		if err != nil {
			return fmt.Errorf("failed to check if module exists: %w", err)
		}
		if exists {
			return fmt.Errorf("module with name '%s' already exists", module.Name)
		}

		attackID, err = s.adpwnModuleRepo.CreateWithObject(ctx, tx, module)
		if err != nil {
			return fmt.Errorf("failed to create adpwn module: %w", err)
		}

		if len(module.Options) != 0 {
			for _, option := range module.Options {
				err := s.adpwnModuleRepo.AddOption(ctx, tx, option)
				if err != nil {
					return fmt.Errorf("error while adding option '%s'", option.Key)
				}
			}
		}

		return nil
	})

	return attackID, err
}

func (s *ADPwnModuleService) CreateModuleInheritanceEdges(ctx context.Context, inheritanceEdges []*adpwn.ModuleInheritanceEdge) error {
	return db.ExecutePostgresInTransaction(ctx, s.db, func(tx *gorm.DB) error {
		log.Printf("CREATE INHERITANCE")
		for _, inheritanceEdge := range inheritanceEdges {
			exists, err := s.adpwnModuleRepo.CheckIfEdgeExits(ctx, tx, inheritanceEdge.PreviousModule, inheritanceEdge.NextModule)
			if err != nil {
				return fmt.Errorf("failed to check if inheritance edge exists: %w", err)
			}

			if exists {
				log.Printf("edge from '%s' to '%s' already exists\n", inheritanceEdge.PreviousModule, inheritanceEdge.NextModule)
				continue
			}

			_, err = s.adpwnModuleRepo.AddInheritanceEdge(ctx, tx, inheritanceEdge.PreviousModule, inheritanceEdge.NextModule)
			if err != nil {
				return fmt.Errorf("failed to add inheritance edge: %w", err)
			}
		}

		return nil
	})
}

func (s *ADPwnModuleService) GetAll(ctx context.Context) ([]*adpwn.Module, error) {
	return s.adpwnModuleRepo.GetAll(ctx, s.db)
}

func (s *ADPwnModuleService) GetInheritanceGraph(ctx context.Context) (*adpwn.InheritanceGraph, error) {
	return db.ExecutePostgresRead(ctx, s.db, func(db *gorm.DB) (*adpwn.InheritanceGraph, error) {
		var inheritanceGraph adpwn.InheritanceGraph

		modules, err := s.adpwnModuleRepo.GetAll(ctx, db)
		if err != nil {
			return nil, fmt.Errorf("error while fetching all adpwn modules %s", err)
		}
		inheritanceGraph.Nodes = modules
		edges, err := s.adpwnModuleRepo.GetAllInheritanceEdges(ctx, db)
		if err != nil {
			return nil, fmt.Errorf("error while fetching all adpwn inheritance edges %s", err)
		}
		inheritanceGraph.Edges = edges
		return &inheritanceGraph, nil
	})
}

func (*ADPwnModuleService) Run(uid string) error {
	panic("Implement me")
}
