package service

import (
	"ADPwn/core/internal/repository"
	"ADPwn/core/internal/utils"
	"ADPwn/core/model"
	"context"
	"github.com/dgraph-io/dgo/v210"
)

type ProjectService struct {
	projectRepo repository.ProjectRepository
	domainRepo  repository.DomainRepository
	hostRepo    repository.HostRepository
	targetRepo  repository.TargetRepository
	DB          *dgo.Dgraph
}

func NewProjectService() (*ProjectService, error) {
	DB, err := utils.GetDB()
	if err != nil {
		return nil, err
	}

	projectRepo := repository.NewDgraphProjectRepository(DB)
	domainRepo := repository.NewDgraphDomainRepository(DB)
	hostRepo := repository.NewDgraphHostRepository(DB)
	targetRepo := repository.NewDgraphTargetRepository(DB)

	return &ProjectService{
		projectRepo: projectRepo,
		domainRepo:  domainRepo,
		hostRepo:    hostRepo,
		targetRepo:  targetRepo}, nil
}

func (s *ProjectService) AddDomainWithHosts(ctx context.Context, projectUID string, domainName string, hosts []string) error {
	tx := s.DB.NewTxn()
	defer tx.Discard(ctx)

	// 1. Create Domain
	domainUID, err := s.domainRepo.Create(ctx, domainName)
	if err != nil {
		return err
	}

	// 2. Connect domain with project
	if err := s.projectRepo.AddDomain(ctx, projectUID, domainUID); err != nil {
		return err
	}

	// Create and connect hosts
	for _, ip := range hosts {
		hostUID, err := s.hostRepo.Create(ctx, ip)
		if err != nil {
			return err
		}

		if err := s.domainRepo.AddHost(ctx, domainUID, hostUID); err != nil {
			return err
		}
	}

	return tx.Commit(ctx)
}

func (s *ProjectService) AddTarget(ctx context.Context, projectUID string, targetIp string) error {
	tx := s.DB.NewTxn()
	defer tx.Discard(ctx)

	// 1. Create target
	targetUID, err := s.targetRepo.Create(ctx, targetIp)
	if err != nil {
		return err
	}

	// 2. Add target to project
	if err := s.projectRepo.AddTarget(ctx, projectUID, targetUID); err != nil {
		return err
	}

	return tx.Commit(ctx)
}

func (s *ProjectService) Create(ctx context.Context, name string) (string, error) {
	return s.projectRepo.Create(ctx, name)
}

func (s *ProjectService) GetOverviewForAll(ctx context.Context) ([]*model.Project, error) {
	return s.projectRepo.GetAllOverview(ctx)
}

func (s *ProjectService) Get(ctx context.Context, projectUID string) (*model.Project, error) {
	return s.projectRepo.Get(ctx, projectUID)
}
