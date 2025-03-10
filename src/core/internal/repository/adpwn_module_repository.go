package repository

import (
	"ADPwn/core/model"
	"context"
	"github.com/dgraph-io/dgo/v210"
)

type ADPwnModuleRepository interface {
	//CRUD
	GetAll(ctx context.Context, tx *dgo.Txn) ([]model.Project, error)
	Create(ctx context.Context, name string) (string, error) // Returns UID
	CreateWithObject(ctx context.Context, module model.ADPwnModule) (string, error)
	Get(ctx context.Context, uid string) (*model.ADPwnModule, error)
	GetByAttackID(ctx context.Context, attackID string) (*model.ADPwnModule, error)
	UpdateFields(ctx context.Context, uid string, fields map[string]interface{}) error

	//Relations
	LinkAsAssumption(ctx context.Context, selfUID, targetUI string) error
	AddDependency(ctx context.Context, selfUID, dependencyUID string) error
}

type DgraphADPwnModuleRepository struct {
	DB *dgo.Dgraph
}

func (r *DgraphADPwnModuleRepository) CreateWithObject(ctx context.Context, module model.ADPwnModule) (string, error) {
	//TODO implement me
	panic("implement me")
}

func (r *DgraphADPwnModuleRepository) Get(ctx context.Context, uid string) (*model.ADPwnModule, error) {
	//TODO implement me
	panic("implement me")
}

func (r *DgraphADPwnModuleRepository) GetByAttackID(ctx context.Context, attackID string) (*model.ADPwnModule, error) {
	//TODO implement me
	panic("implement me")
}

func (r *DgraphADPwnModuleRepository) UpdateFields(ctx context.Context, uid string, fields map[string]interface{}) error {
	//TODO implement me
	panic("implement me")
}

func (r *DgraphADPwnModuleRepository) LinkAsAssumption(ctx context.Context, selfUID, targetUI string) error {
	//TODO implement me
	panic("implement me")
}

func (r *DgraphADPwnModuleRepository) AddDependency(ctx context.Context, selfUID, dependencyUID string) error {
	//TODO implement me
	panic("implement me")
}

func (r *DgraphADPwnModuleRepository) Create(ctx context.Context, name string) (string, error) {
	//TODO implement me
	panic("implement me")
}
