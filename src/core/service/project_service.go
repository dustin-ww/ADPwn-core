package service

import (
	"ADPwn/core/internal/db"
	"ADPwn/core/internal/repository"
	"ADPwn/core/internal/utils"
	"ADPwn/core/model"
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"github.com/dgraph-io/dgo/v210/protos/api"

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
//func (s *ProjectService) AddDomainWithHosts(ctx context.Context, projectUID, domainName string, hosts []string) error {
//	return db.ExecuteInTransaction(ctx, s.db, func(tx *dgo.Txn) error {
//		domainUID, err := s.domainRepo.Create(ctx, domainName)
//		if err != nil {
//			return fmt.Errorf("failed to create domain: %w", err)
//		}
//
//		if err := s.projectRepo.AddDomain(ctx, tx, projectUID, domainUID); err != nil {
//			return fmt.Errorf("failed to link domain: %w", err)
//		}
//
//		for _, ip := range hosts {
//			hostUID, err := s.hostRepo.Create(ctx, ip)
//			if err != nil {
//				return fmt.Errorf("failed to create host %s: %w", ip, err)
//			}
//			if err := s.domainRepo.AddHost(ctx, domainUID, hostUID); err != nil {
//				return fmt.Errorf("failed to link host: %w", err)
//			}
//		}
//		return nil
//	})
//}

func (s *ProjectService) AddDomain(ctx context.Context, projectUID string, domain *model.Domain) error {
	return db.ExecuteInTransaction(ctx, s.db, func(tx *dgo.Txn) error {
		domainUID, err := s.domainRepo.CreateWithObject(ctx, tx, domain)
		if err != nil {
			return fmt.Errorf("failed to create domain: %w", err)
		}

		if err := s.projectRepo.AddDomain(ctx, tx, projectUID, domainUID); err != nil {
			return fmt.Errorf("failed to link domain: %w", err)
		}

		if err := s.domainRepo.AddToProject(ctx, tx, domainUID, projectUID); err != nil {
			return fmt.Errorf("failed to reverse link domain to project: %w", err)
		}

		//if len(domain.HasHost) > 0 {
		//	for _, host := range domain.HasHost {
		//		hostUID, err := s.hostRepo.Create(ctx, tx, host.Name)
		//		if err != nil {
		//			return fmt.Errorf("failed to create host %s: %w", host.Name, err)
		//		}
		//		if err := s.domainRepo.AddHost(ctx, tx, domainUID, hostUID); err != nil {
		//			return fmt.Errorf("failed to link host: %w", err)
		//		}
		//	}
		//}
		return nil
	})
}

func (s *ProjectService) GetProjectDomains(ctx context.Context, projectUID string) ([]*model.Domain, error) {
	return db.ExecuteRead(ctx, s.db, func(tx *dgo.Txn) ([]*model.Domain, error) {
		return s.domainRepo.GetByProjectUID(ctx, tx, projectUID)
	})
}

// CreateTarget creates a new target and links it to a project.
func (s *ProjectService) CreateTargets(ctx context.Context, projectUID, ip, note string, cidr int) ([]string, error) {
	ips, err := utils.IpsFromIPAndCIDR(ip, cidr)
	if err != nil {
		return nil, fmt.Errorf("failed to get ips: %w", err)
	}

	// Batch-Größe definieren (z.B. 1000 IPs pro Batch)
	batchSize := 1000
	allTargetUIDs := make([]string, 0, len(ips))

	err = db.ExecuteInTransaction(ctx, s.db, func(tx *dgo.Txn) error {
		for i := 0; i < len(ips); i += batchSize {
			end := i + batchSize
			if end > len(ips) {
				end = len(ips)
			}
			batch := ips[i:end]

			// 1. Massen-Mutation für Targets
			targetUIDs, err := s.createTargetsBatch(ctx, tx, batch, note)
			if err != nil {
				return err
			}
			allTargetUIDs = append(allTargetUIDs, targetUIDs...)

			// 2. Massen-Verknüpfung mit Projekt
			if err := s.linkTargetsToProject(ctx, tx, projectUID, targetUIDs); err != nil {
				return err
			}
		}
		return nil
	})

	return allTargetUIDs, err
}

func (s *ProjectService) createTargetsBatch(ctx context.Context, tx *dgo.Txn, ips []string, note string) ([]string, error) {
	targets := make([]interface{}, len(ips))
	for i, ip := range ips {
		targets[i] = map[string]interface{}{
			"uid":         fmt.Sprintf("_:target%d", i), // Unique Blank Node
			"ip":          ip,
			"note":        note,
			"dgraph.type": "Target",
		}
	}

	targetJSON, _ := json.Marshal(targets)
	mu := &api.Mutation{SetJson: targetJSON}
	assigned, err := tx.Mutate(ctx, mu)
	if err != nil {
		return nil, fmt.Errorf("batch mutation failed: %w", err)
	}

	uids := make([]string, len(ips))
	for i := 0; i < len(ips); i++ {
		uidKey := fmt.Sprintf("target%d", i)
		uids[i] = assigned.Uids[uidKey]
	}
	return uids, nil
}

func (s *ProjectService) linkTargetsToProject(ctx context.Context, tx *dgo.Txn, projectUID string, targetUIDs []string) error {
	nquads := bytes.Buffer{}
	for _, uid := range targetUIDs {
		nquads.WriteString(fmt.Sprintf("<%s> <has_target> <%s> .\n", projectUID, uid))
	}

	mu := &api.Mutation{
		SetNquads: nquads.Bytes(),
	}
	_, err := tx.Mutate(ctx, mu)
	return err
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
		return s.projectRepo.GetAll(ctx, tx)
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
