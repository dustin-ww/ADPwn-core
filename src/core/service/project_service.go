package service

import (
	"ADPwn/core/internal/db"
	"ADPwn/core/internal/utils"
	"ADPwn/core/model"
	"context"
	"fmt"
	"log"

	"ADPwn/core/internal/repository"

	"github.com/dgraph-io/dgo/v210"
)

// ProjectService handles business logic for projects
type ProjectService struct {
	projectRepo repository.ProjectRepository
	domainRepo  repository.DomainRepository
	hostRepo    repository.HostRepository
	targetRepo  repository.TargetRepository
	db          *dgo.Dgraph
}

// NewProjectService creates a new ProjectService instance
func NewProjectService() (*ProjectService, error) {
	db, err := db.GetDB()
	if err != nil {
		return nil, fmt.Errorf("failed to get database connection: %w", err)
	}

	return &ProjectService{
		db:          db,
		projectRepo: repository.NewDgraphProjectRepository(),
		domainRepo:  repository.NewDgraphDomainRepository(db),
		hostRepo:    repository.NewDgraphHostRepository(db),
		targetRepo:  repository.NewDgraphTargetRepository(db),
	}, nil
}

// AddDomainWithHosts adds a domain with associated hosts to a project.
func (s *ProjectService) AddDomainWithHosts(ctx context.Context, projectUID, domainName string, hosts []string) error {
	return db.ExecuteInTransaction(ctx, s.db, func(tx *dgo.Txn) error {
		domainUID, err := s.domainRepo.Create(ctx, domainName)
		if err != nil {
			return fmt.Errorf("failed to create domain: %w", err)
		}

		if err := s.projectRepo.AddDomain(ctx, tx, projectUID, domainUID); err != nil {
			return fmt.Errorf("failed to link domain: %w", err)
		}

		for _, ip := range hosts {
			hostUID, err := s.hostRepo.Create(ctx, ip)
			if err != nil {
				return fmt.Errorf("failed to create host %s: %w", ip, err)
			}
			if err := s.domainRepo.AddHost(ctx, domainUID, hostUID); err != nil {
				return fmt.Errorf("failed to link host: %w", err)
			}
		}
		return nil
	})
}

// CreateTarget creates a new target and links it to a project.
func (s *ProjectService) CreateTarget(ctx context.Context, projectUID, targetIP, name string) (string, error) {
	var targetUID string
	err := db.ExecuteInTransaction(ctx, s.db, func(tx *dgo.Txn) error {
		var err error
		targetUID, err = s.targetRepo.Create(ctx, tx, targetIP, name)
		if err != nil {
			return fmt.Errorf("target creation failed: %w", err)
		}
		log.Println("Target created:", targetUID)

		if err := s.projectRepo.AddTarget(ctx, tx, projectUID, targetUID); err != nil {
			return fmt.Errorf("linking failed: %w", err)
		}
		return nil
	})
	return targetUID, err
}

// Create creates a new project.
func (s *ProjectService) Create(ctx context.Context, name string) (string, error) {
	var projectUID string
	err := db.ExecuteInTransaction(ctx, s.db, func(tx *dgo.Txn) error {
		var err error
		projectUID, err = s.projectRepo.Create(ctx, tx, name)
		if err != nil {
			return fmt.Errorf("failed to create project: %w", err)
		}
		return nil
	})
	return projectUID, err
}

// GetOverviewForAll retrieves overview information for all projects.
func (s *ProjectService) GetOverviewForAll(ctx context.Context) ([]*model.Project, error) {
	return db.ExecuteRead(ctx, s.db, func(tx *dgo.Txn) ([]*model.Project, error) {
		return s.projectRepo.GetAllOverview(ctx, tx)
	})
}

// Get retrieves a project by its UID.
func (s *ProjectService) Get(ctx context.Context, projectUID string) (*model.Project, error) {
	return db.ExecuteRead(ctx, s.db, func(tx *dgo.Txn) (*model.Project, error) {
		return s.projectRepo.Get(ctx, tx, projectUID)
	})
}

// UpdateFields updates specified fields of a project.
func (s *ProjectService) UpdateFields(ctx context.Context, uid string, fields map[string]interface{}) error {
	if uid == "" {
		return utils.ErrUIDRequired
	}

	allowed := map[string]bool{"name": true, "description": true}
	protected := map[string]bool{"uid": true, "created_at": true, "updated_at": true, "type": true}

	for field := range fields {
		if protected[field] {
			return fmt.Errorf("%w: %s", utils.ErrFieldProtected, field)
		}
		if !allowed[field] {
			return fmt.Errorf("%w: %s", utils.ErrFieldNotAllowed, field)
		}
	}

	return db.ExecuteInTransaction(ctx, s.db, func(tx *dgo.Txn) error {
		return s.projectRepo.UpdateFields(ctx, tx, uid, fields)
	})
}

// GetTargets retrieves all targets associated with a project.
func (s *ProjectService) GetTargets(ctx context.Context, projectUID string) ([]*model.Target, error) {
	return db.ExecuteRead(ctx, s.db, func(tx *dgo.Txn) ([]*model.Target, error) {
		return s.projectRepo.GetTargets(ctx, tx, projectUID)
	})
}
