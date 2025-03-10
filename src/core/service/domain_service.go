package service

import (
	"ADPwn/core/internal/db"
	"ADPwn/core/internal/repository"
	"github.com/dgraph-io/dgo/v210"
)

type DomainService struct {
	domainRepo repository.DomainRepository
	DB         *dgo.Dgraph
}

func NewDomainService() (*DomainService, error) {
	DB, err := db.GetDB()
	if err != nil {
		return nil, err
	}

	domainRepo := repository.NewDgraphDomainRepository(DB)

	return &DomainService{
		domainRepo: domainRepo,
	}, nil
}
