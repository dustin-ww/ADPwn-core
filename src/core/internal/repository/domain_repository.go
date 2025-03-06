package repository

import (
	"ADPwn/core/model"
	"context"
	"fmt"
	"github.com/dgraph-io/dgo/v210"
	"github.com/dgraph-io/dgo/v210/protos/api"
)

type DomainRepository interface {
	//CRUD
	Create(ctx context.Context, name string) (string, error) // Returns UID
	Get(ctx context.Context, uid string) (*model.Domain, error)
	UpdateFields(ctx context.Context, uid string, fields map[string]interface{}) error

	//Relations
	LinkToProject(ctx context.Context, domainUID, projectUID string) error
	AddHost(ctx context.Context, domainUID, hostUID string) error
	AddUser(ctx context.Context, domainUID, userUID string) error
}

type DgraphDomainRepository struct {
	DB *dgo.Dgraph
}

func (r *DgraphDomainRepository) Create(ctx context.Context, name string) (string, error) {
	//TODO implement me
	panic("implement me")
}

func (r *DgraphDomainRepository) Get(ctx context.Context, uid string) (*model.Domain, error) {
	//TODO implement me
	panic("implement me")
}

func (r *DgraphDomainRepository) UpdateFields(ctx context.Context, uid string, fields map[string]interface{}) error {
	//TODO implement me
	panic("implement me")
}

func (r *DgraphDomainRepository) AddHost(ctx context.Context, domainUID, hostUID string) error {
	//TODO implement me
	panic("implement me")
}

func (r *DgraphDomainRepository) AddUser(ctx context.Context, domainUID, userUID string) error {
	//TODO implement me
	panic("implement me")
}

func NewDgraphDomainRepository(db *dgo.Dgraph) *DgraphDomainRepository {
	return &DgraphDomainRepository{DB: db}
}

func (r *DgraphDomainRepository) LinkToProject(ctx context.Context, domainUID, projectUID string) error {
	mu := &api.Mutation{
		SetNquads: []byte(fmt.Sprintf(
			`<%s> <has_domain> <%s> .`, // Project → Domain
			projectUID, domainUID,
		)),
	}
	_, err := r.DB.NewTxn().Mutate(ctx, mu)
	return err
}
