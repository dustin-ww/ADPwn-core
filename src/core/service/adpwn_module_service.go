package service

import (
	"ADPwn/core/interfaces"
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
	attackRunner    interfaces.ModuleExecutor // Add this back
}

func NewADPwnModuleService(attackRunner interfaces.ModuleExecutor) (*ADPwnModuleService, error) {
	db, err := db.GetPostgresDB()
	if err != nil {
		return nil, fmt.Errorf("failed to get database connection: %w", err)
	}

	return &ADPwnModuleService{
		db:              db,
		adpwnModuleRepo: repository.NewPostgresADPwnModuleRepository(),
		attackRunner:    attackRunner, // Store the executor
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

func (s *ADPwnModuleService) CreateModuleInheritanceEdges(ctx context.Context, inheritanceEdges []*adpwn.ModuleDependency) error {
	return db.ExecutePostgresInTransaction(ctx, s.db, func(tx *gorm.DB) error {
		for _, inheritanceEdge := range inheritanceEdges {
			exists, err := s.adpwnModuleRepo.CheckIfDependencyExits(ctx, tx, inheritanceEdge.PreviousModule, inheritanceEdge.NextModule)
			if err != nil {
				return fmt.Errorf("failed to check if inheritance edge exists: %w", err)
			}

			if exists {
				log.Printf("edge from '%s' to '%s' already exists\n", inheritanceEdge.PreviousModule, inheritanceEdge.NextModule)
				continue
			}

			_, err = s.adpwnModuleRepo.AddDependency(ctx, tx, inheritanceEdge.PreviousModule, inheritanceEdge.NextModule)
			if err != nil {
				return fmt.Errorf("failed to add inheritance edge: %w", err)
			}
		}

		return nil
	})
}

func (s *ADPwnModuleService) GetAll(ctx context.Context) ([]*adpwn.Module, error) {
	return db.ExecutePostgresRead(ctx, s.db, func(db *gorm.DB) ([]*adpwn.Module, error) {
		modules, err := s.adpwnModuleRepo.GetAll(ctx, db)
		if err != nil {
			return nil, err
		}

		for _, module := range modules {
			module.Options, err = s.adpwnModuleRepo.GetOptions(ctx, db, module.Key)
			if err != nil {
				return nil, fmt.Errorf("failed to get module options: %w", err)
			}
			module.DependencyVector, err = s.adpwnModuleRepo.GetOrderedDependencies(ctx, db, module.Key)
			if err != nil {
				return nil, fmt.Errorf("failed to get module dependency vector: %w", err)
			}
		}

		return modules, nil
	})
}

func (s *ADPwnModuleService) GetAttackVectorByKey(ctx context.Context, moduleKey string) ([]*adpwn.Module, error) {
	return db.ExecutePostgresRead(ctx, s.db, func(db *gorm.DB) ([]*adpwn.Module, error) {
		var attackVector []*adpwn.Module

		module, err := s.adpwnModuleRepo.Get(ctx, db, moduleKey)
		if err != nil {
			return nil, fmt.Errorf("error while fetching all dependencies of adpwn modules for graph edges %s", err)
		}

		if module.DependencyVector == nil {
			module.DependencyVector, err = s.adpwnModuleRepo.GetOrderedDependencies(ctx, db, moduleKey)
		}

		for _, dependencyKey := range module.DependencyVector {
			dep, err := s.adpwnModuleRepo.Get(ctx, db, dependencyKey)
			if err != nil {
				return nil, fmt.Errorf("error while fetching all dependencies of adpwn modules for graph edges %s", err)
			}
			attackVector = append(attackVector, dep)
		}
		return append(attackVector, module), nil
	})
}

//func (s *ADPwnModuleService) GetAttackVector(ctx context.Context, moduleKey string) ([]*adpwn.Module, error) {
//	return db.ExecutePostgresRead(ctx, s.db, func(db *gorm.DB) (*adpwn.Module, error) {
//		s.adpwnModuleRepo.GetOrderedDependencies()
//	})
//}

func (s *ADPwnModuleService) GetInheritanceGraph(ctx context.Context) (*adpwn.InheritanceGraph, error) {
	return db.ExecutePostgresRead(ctx, s.db, func(db *gorm.DB) (*adpwn.InheritanceGraph, error) {
		var inheritanceGraph adpwn.InheritanceGraph

		modules, err := s.adpwnModuleRepo.GetAll(ctx, db)
		if err != nil {
			return nil, fmt.Errorf("error while fetching all adpwn modules %s", err)
		}
		inheritanceGraph.Nodes = modules
		edges, err := s.adpwnModuleRepo.GetAllDependencies(ctx, db)
		if err != nil {
			return nil, fmt.Errorf("error while fetching all dependencies of adpwn modules for graph edges %s", err)
		}
		inheritanceGraph.Edges = edges
		return &inheritanceGraph, nil
	})
}

func (s *ADPwnModuleService) RunAttackVector(ctx context.Context, key string) error {
	return db.ExecutePostgresInTransaction(ctx, s.db, func(tx *gorm.DB) error {
		// Use the attackRunner that was injected into the service
		err := RunAttackVector(ctx, key, nil, s.attackRunner)
		if err != nil {
			return err
		}
		return nil
	})
}

func (s *ADPwnModuleService) GetOptionsForAttackVector(ctx context.Context, moduleKey string) ([]*adpwn.ModuleOption, error) {
	return db.ExecutePostgresRead(ctx, s.db, func(db *gorm.DB) ([]*adpwn.ModuleOption, error) {
		modules, err := s.GetAttackVectorByKey(ctx, moduleKey)
		log.Println(moduleKey)
		log.Println(len(modules))
		if err != nil {
			return nil, err
		}

		for _, module := range modules {
			module.Options, err = s.adpwnModuleRepo.GetOptions(ctx, db, module.Key)
			if err != nil {
				return nil, fmt.Errorf("failed to get module options: %w", err)
			}
		}

		seenKeys := make(map[string]struct{})
		uniqueOptions := make([]*adpwn.ModuleOption, 0)

		for _, module := range modules {
			for _, option := range module.Options {
				log.Println(option.Key)
				key := option.Key
				if _, exists := seenKeys[key]; !exists {
					seenKeys[key] = struct{}{}
					uniqueOptions = append(uniqueOptions, option)
				}
			}
		}
		return uniqueOptions, nil
	})

}
