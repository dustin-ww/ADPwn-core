package service

import (
	"ADPwn-core/internal/db"
	"ADPwn-core/internal/repository"
	"ADPwn-core/pkg/model"
	"context"
	"github.com/dgraph-io/dgo/v210"
)

type HostService struct {
	hostRepo    repository.HostRepository
	serviceRepo repository.ServiceRepository
	DB          *dgo.Dgraph
}

func NewHostService() (*HostService, error) {
	DB, err := db.GetDB()
	if err != nil {
		return nil, err
	}

	hostRepo := repository.NewDgraphHostRepository(DB)
	serviceRepo := repository.NewDgraphServiceRepository(DB)

	return &HostService{
		hostRepo:    hostRepo,
		serviceRepo: serviceRepo}, nil
}

func (s *HostService) AddService(ctx context.Context, hostUID string, service model.Service) error {

	// 1. Create Service
	serviceUID, err := s.serviceRepo.CreateWithObject(ctx, service)
	if err != nil {
		return err
	}

	// 2. Connect service with host
	if err := s.hostRepo.AddService(ctx, hostUID, serviceUID); err != nil {
		return err
	}

	return nil

}
