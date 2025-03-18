package repository

import (
	"ADPwn/core/model/adpwn"
	"context"
	"fmt"
	"gorm.io/gorm"
	"log"
)

type ADPwnCollectionRepository interface {
	//CRUD
	GetAll(ctx context.Context, tx *gorm.DB) ([]*adpwn.Collection, error)
	Create(ctx context.Context, tx *gorm.DB, name string, description string) (uint, error)
	AddModule(ctx context.Context, tx *gorm.DB, collectionID, moduleKey string)
	GetModulesForCollection(ctx context.Context, tx *gorm.DB, collectionID uint) ([]adpwn.Module, error)
}

type PostgresADPwnCollectionRepository struct {
}

func NewPostgresADPwnCollectionRepository() *PostgresADPwnCollectionRepository {
	return &PostgresADPwnCollectionRepository{}
}

func (r *PostgresADPwnCollectionRepository) GetAll(ctx context.Context, tx *gorm.DB) ([]*adpwn.Collection, error) {
	var collections []*adpwn.Collection

	err := tx.WithContext(ctx).Table("adpwn_collections").Find(&collections).Error
	if err != nil {
		return nil, fmt.Errorf("failed to get all adpwn collections: %w", err)
	}

	return collections, nil
}

func (r *PostgresADPwnCollectionRepository) GetModulesForCollection(ctx context.Context, tx *gorm.DB, collectionID uint) ([]adpwn.Module, error) {
	var moduleKeys []string

	err := tx.WithContext(ctx).
		Table("adpwn_collection_modules").
		Where("collection_id = ?", collectionID).
		Pluck("module_key", &moduleKeys).
		Error

	if err != nil {
		return nil, fmt.Errorf("failed to get module keys for collection %d: %w", collectionID, err)
	}

	if len(moduleKeys) == 0 {
		return []adpwn.Module{}, nil
	}

	// Module laden
	var modules []adpwn.Module

	err = tx.WithContext(ctx).
		Table("adpwn_modules").
		Where("key IN ?", moduleKeys).
		Find(&modules).
		Error

	if err != nil {
		return nil, fmt.Errorf("failed to get modules for collection %d: %w", collectionID, err)
	}

	return modules, nil
}

func (r *PostgresADPwnCollectionRepository) Create(ctx context.Context, tx *gorm.DB, name string, description string) (uint, error) {
	collection := &adpwn.Collection{Name: name, Description: description}

	err := tx.WithContext(ctx).Table("adpwn_collections").Create(collection).Error
	if err != nil {
		log.Printf("failed to create adpwn collection: %v", err)
		return 0, err
	}

	return collection.ID, nil
}

func (r *PostgresADPwnCollectionRepository) AddModule(ctx context.Context, tx *gorm.DB, collectionID, moduleKey string) {

}
