package repository

import (
	"ADPwn-core/pkg/model/adpwn"
	"context"
	"errors"
	"fmt"
	"gorm.io/gorm"
	"log"
)

const (
	TableModules            = "adpwn_modules"
	TableModuleDependencies = "adpwn_modules_dependencies"
	TableModuleOptions      = "adpwn_modules_options"
)

// ADPwnModuleRepository This is a repository for implementing Postgres-specific persistence of adpwn modulelib.
type ADPwnModuleRepository interface {

	//CRUD
	GetAll(ctx context.Context, tx *gorm.DB) ([]*adpwn.Module, error)
	CreateWithObject(ctx context.Context, tx *gorm.DB, module *adpwn.Module) (string, error)
	Get(ctx context.Context, tx *gorm.DB, moduleKey string) (*adpwn.Module, error)
	CheckIfExistsByKey(ctx context.Context, tx *gorm.DB, key string) (bool, error)

	// module dependencies
	CheckIfDependencyExits(ctx context.Context, tx *gorm.DB, previousModuleKey, nextModuleKey string) (bool, error)
	AddDependency(ctx context.Context, tx *gorm.DB, previousModuleKey, nextModuleKey string) (string, error)
	GetAllDependencies(ctx context.Context, tx *gorm.DB) ([]*adpwn.ModuleDependency, error)
	GetOrderedDependencies(ctx context.Context, tx *gorm.DB, moduleKey string) ([]string, error)

	// module options
	AddOption(ctx context.Context, tx *gorm.DB, moduleOption *adpwn.ModuleOption) error
	GetOptions(ctx context.Context, tx *gorm.DB, moduleKey string) ([]*adpwn.ModuleOption, error)
}

type PostgresADPwnModuleRepository struct{}

func (r *PostgresADPwnModuleRepository) GetOrderedDependencies(ctx context.Context, tx *gorm.DB, moduleKey string) ([]string, error) {
	// SQL statement for recursive Common Table Expression (CTE)
	// This mimics the WITH RECURSIVE functionality from PostgreSQL
	// See: https://www.dylanpaulus.com/posts/postgres-is-a-graph-database/
	query := `
        WITH RECURSIVE dependent_modules AS (
            SELECT previous_module
            FROM adpwn_modules_dependencies
            WHERE next_module = ?
            
            UNION
            
            SELECT e.previous_module
            FROM adpwn_modules_dependencies e
            JOIN dependent_modules dm ON e.next_module = dm.previous_module
        )
        SELECT m.key
        FROM adpwn_modules m
        JOIN dependent_modules dm ON m.key = dm.previous_module
    `

	var moduleKeys []string

	if err := tx.WithContext(ctx).Raw(query, moduleKey).Scan(&moduleKeys).Error; err != nil {
		return nil, fmt.Errorf("failed to get dependency key list: %w", err)
	}

	return moduleKeys, nil
}

func (r *PostgresADPwnModuleRepository) AddOption(ctx context.Context, tx *gorm.DB, moduleOption *adpwn.ModuleOption) error {
	result := tx.WithContext(ctx).Table(TableModuleOptions).Create(&moduleOption)
	if result.Error != nil {
		return fmt.Errorf("create failed: %w", result.Error)
	}
	return nil
}

func (r *PostgresADPwnModuleRepository) AddDependency(ctx context.Context, tx *gorm.DB, previousModuleKey, nextModuleKey string) (string, error) {
	dependency := &adpwn.ModuleDependency{PreviousModule: previousModuleKey, NextModule: nextModuleKey}
	result := tx.WithContext(ctx).Table(TableModuleDependencies).Create(&dependency)
	if result.Error != nil {
		return "", fmt.Errorf("create failed: %w", result.Error)
	}
	//TODO: Change
	return dependency.PreviousModule, nil
}

func (r *PostgresADPwnModuleRepository) GetAllDependencies(ctx context.Context, tx *gorm.DB) ([]*adpwn.ModuleDependency, error) {
	var dependencies []*adpwn.ModuleDependency

	err := tx.WithContext(ctx).Table(TableModuleDependencies).Find(&dependencies).Error

	if err != nil {
		return nil, fmt.Errorf("failed to get all edges: %w", err)
	}

	if len(dependencies) == 0 {
		log.Println("no edges found in database")
	}
	return dependencies, nil
}

func (r *PostgresADPwnModuleRepository) GetAll(ctx context.Context, tx *gorm.DB) ([]*adpwn.Module, error) {
	var modules []*adpwn.Module

	err := tx.WithContext(ctx).Table(TableModules).Find(&modules).Error

	if err != nil {
		return nil, fmt.Errorf("failed to get all modulelib: %w", err)
	}

	if len(modules) == 0 {
		log.Println("no modulelib found in database")
	}
	return modules, nil
}

func NewPostgresADPwnModuleRepository() *PostgresADPwnModuleRepository {
	return &PostgresADPwnModuleRepository{}
}

func (r *PostgresADPwnModuleRepository) CheckIfDependencyExits(ctx context.Context, tx *gorm.DB, previousModuleKey, nextModuleKey string) (bool, error) {
	query := tx.WithContext(ctx)
	var count int64
	err := query.Table(TableModuleDependencies).
		Where("previous_module = ?", previousModuleKey).
		Where("next_module = ?", nextModuleKey).
		Count(&count).
		Error
	if err != nil {
		return false, fmt.Errorf("failed to check if module exists by key: %w", err)
	}
	return count > 0, nil
}

func (r *PostgresADPwnModuleRepository) CheckIfExistsByKey(ctx context.Context, tx *gorm.DB, key string) (bool, error) {
	query := tx.WithContext(ctx)
	var count int64
	err := query.Table(TableModules).
		Where("key = ?", key).
		Count(&count).
		Error
	if err != nil {
		return false, fmt.Errorf("failed to check if module exists by key: %w", err)
	}
	return count > 0, nil
}

func (r *PostgresADPwnModuleRepository) CreateWithObject(ctx context.Context, tx *gorm.DB, module *adpwn.Module) (string, error) {
	result := tx.WithContext(ctx).Table(TableModules).Create(&module)
	if result.Error != nil {
		return "", fmt.Errorf("create failed: %w", result.Error)
	}

	return module.AttackID, nil
}

func (r *PostgresADPwnModuleRepository) Get(ctx context.Context, tx *gorm.DB, moduleKey string) (*adpwn.Module, error) {
	{
		var module adpwn.Module

		tx := tx.WithContext(ctx)

		err := tx.Table(TableModules).
			First(&module, "key = ?", moduleKey).
			Error

		if err != nil {
			if errors.Is(err, gorm.ErrRecordNotFound) {
				return nil, fmt.Errorf("module not found: %s", moduleKey)
			}
			return nil, fmt.Errorf("database error: %w", err)
		}
		return &module, nil
	}
}

func (r *PostgresADPwnModuleRepository) GetOptions(ctx context.Context, tx *gorm.DB, moduleKey string) ([]*adpwn.ModuleOption, error) {
	if moduleKey == "" {
		return nil, errors.New("moduleKey cannot be empty")
	}

	var options []*adpwn.ModuleOption

	result := tx.WithContext(ctx).
		Table(TableModuleOptions).
		Where("module_key = ?", moduleKey).
		Find(&options)

	if err := result.Error; err != nil {
		return nil, fmt.Errorf("failed to fetch module options: %w", err)
	}

	if options == nil {
		options = []*adpwn.ModuleOption{}
	}

	return options, nil
}
