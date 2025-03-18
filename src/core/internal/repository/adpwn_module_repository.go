package repository

import (
	"ADPwn/core/model/adpwn"
	"context"
	"errors"
	"fmt"
	"gorm.io/gorm"
	"log"
)

type ADPwnModuleRepository interface {
	//CRUD
	GetAll(ctx context.Context, tx *gorm.DB) ([]*adpwn.Module, error)
	CreateWithObject(ctx context.Context, tx *gorm.DB, module *adpwn.Module) (string, error)
	Get(ctx context.Context, tx *gorm.DB, attack_id string) (*adpwn.Module, error)
	CheckIfExistsByKey(ctx context.Context, tx *gorm.DB, key string) (bool, error)

	// Inheritance Edges
	CheckIfEdgeExits(ctx context.Context, tx *gorm.DB, previousModuleKey, nextModuleKey string) (bool, error)
	AddInheritanceEdge(ctx context.Context, tx *gorm.DB, previousModuleKey, nextModuleKey string) (string, error)
	GetAllInheritanceEdges(ctx context.Context, tx *gorm.DB) ([]*adpwn.ModuleInheritanceEdge, error)
}

type PostgresADPwnModuleRepository struct{}

func (r *PostgresADPwnModuleRepository) AddInheritanceEdge(ctx context.Context, tx *gorm.DB, previousModuleKey, nextModuleKey string) (string, error) {
	inheritanceEdge := &adpwn.ModuleInheritanceEdge{PreviousModule: previousModuleKey, NextModule: nextModuleKey}
	result := tx.WithContext(ctx).Table("adpwn_modules_edges").Create(&inheritanceEdge)
	if result.Error != nil {
		return "", fmt.Errorf("create failed: %w", result.Error)
	}
	//TODO: Change
	return inheritanceEdge.PreviousModule, nil
}

func (r *PostgresADPwnModuleRepository) GetAllInheritanceEdges(ctx context.Context, tx *gorm.DB) ([]*adpwn.ModuleInheritanceEdge, error) {
	var inheritanceEdges []*adpwn.ModuleInheritanceEdge

	err := tx.WithContext(ctx).Table("adpwn_modules_edges").Find(&inheritanceEdges).Error

	if err != nil {
		return nil, fmt.Errorf("failed to get all edges: %w", err)
	}

	if len(inheritanceEdges) == 0 {
		log.Println("no edges found in database")
	}
	return inheritanceEdges, nil
}

func (r *PostgresADPwnModuleRepository) GetAll(ctx context.Context, tx *gorm.DB) ([]*adpwn.Module, error) {
	var modules []*adpwn.Module

	err := tx.WithContext(ctx).Table("adpwn_modules").Find(&modules).Error

	if err != nil {
		return nil, fmt.Errorf("failed to get all modules: %w", err)
	}

	if len(modules) == 0 {
		log.Println("no modules found in database")
	}
	return modules, nil
}

func NewPostgresADPwnModuleRepository() *PostgresADPwnModuleRepository {
	return &PostgresADPwnModuleRepository{}
}

func (r *PostgresADPwnModuleRepository) AddDependencyEdge(ctx context.Context, tx *gorm.DB, previousModuleKey, nextModuleKey string) error {
	panic("implement me")
}

func (r *PostgresADPwnModuleRepository) CheckIfEdgeExits(ctx context.Context, tx *gorm.DB, previousModuleKey, nextModuleKey string) (bool, error) {
	query := tx.WithContext(ctx)
	var count int64
	err := query.Table("adpwn_modules_edges").
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
	err := query.Table("adpwn_modules").
		Where("key = ?", key).
		Count(&count).
		Error
	if err != nil {
		return false, fmt.Errorf("failed to check if module exists by key: %w", err)
	}
	return count > 0, nil
}

func (r *PostgresADPwnModuleRepository) CreateWithObject(ctx context.Context, tx *gorm.DB, module *adpwn.Module) (string, error) {
	result := tx.WithContext(ctx).Table("adpwn_modules").Create(&module)
	if result.Error != nil {
		return "", fmt.Errorf("create failed: %w", result.Error)
	}

	return module.AttackID, nil
}

func (r *PostgresADPwnModuleRepository) Get(ctx context.Context, tx *gorm.DB, attack_id string) (*adpwn.Module, error) {
	{
		var module adpwn.Module

		tx := tx.WithContext(ctx)

		err := tx.
			Preload("Dependencies").              // Eager Loading von Abhängigkeiten
			First(&module, "uid = ?", attack_id). // UID ist kein Primärschlüssel
			Error

		if err != nil {
			if errors.Is(err, gorm.ErrRecordNotFound) {
				return nil, fmt.Errorf("module not found: %s", attack_id)
			}
			return nil, fmt.Errorf("database error: %w", err)
		}
		return &module, nil
	}
}
