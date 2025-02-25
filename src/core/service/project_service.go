package service

import (
	"ADPwn/core/internal/repository"
	"ADPwn/core/internal/utils"
	"ADPwn/core/model"
	"context"
)

type ProjectService struct {
	repo repository.ProjectRepository
}

func NewProjectService() (*ProjectService, error) {
	db, err := utils.GetDB()
	if err != nil {
		return nil, err
	}

	projectRepo := repository.NewDgraphIOProjectRepository(db)
	return &ProjectService{repo: projectRepo}, nil
}

func (s *ProjectService) AllProjects(ctx context.Context) ([]model.Project, error) {
	return s.repo.AllProjects(ctx)
}

func (s *ProjectService) CreateProject(ctx context.Context, name string) (model.Project, error) {
	return s.repo.SaveProject(ctx, *model.NewProject(name))
}

func (s *ProjectService) SaveProject(ctx context.Context, project model.Project) (model.Project, error) {
	return s.repo.SaveProject(ctx, project)
}

func (s *ProjectService) SaveSubnetTarget(ctx context.Context, project model.Project, subnet string) (model.Project, error) {
	/*IPs, err := utils.GenerateIPs(subnet)

	    if err != nil {
			return model.Project{}, fmt.Errorf("error in project service while generating hosts from subnetmask: %w", err)
		}

		var hosts []string

		for _, ip := range IPs {
			hosts = append(hosts, ip)
		}*/

	project.Targets = append(project.Targets, subnet)

	return s.repo.SaveProject(ctx, project)
}

func (s *ProjectService) SaveSingleTarget(ctx context.Context, project model.Project, ip string) (model.Project, error) {
	project.Targets = append(project.Targets, ip)
	return s.repo.SaveProject(ctx, project)
}

func (s *ProjectService) CreateHost(ctx context.Context, project model.Project, ip string) (model.Project, error) {

	host := *model.NewHost(ip, project.UID, project.Name)
	project.Domains[0].Hosts = append(project.Domains[0].Hosts, host)

	return s.repo.SaveProject(ctx, project)
}

func (s *ProjectService) AddHost(ctx context.Context, project model.Project, host model.Host) (model.Project, error) {

	project.Domains[0].Hosts = append(project.Domains[0].Hosts, host)
	return s.repo.SaveProject(ctx, project)
}

func (s *ProjectService) SaveUser(ctx context.Context, project model.Project, username string, password string, NTLMHash string, isAdmin bool) (model.Project, error) {

	user := *model.NewUser(username, password, NTLMHash, isAdmin)

	project.Domains[0].Users = append(project.Domains[0].Users, user)

	return s.repo.SaveProject(ctx, project)
}

func (s *ProjectService) DeleteProject(ctx context.Context, project model.Project) error {
	return s.repo.DeleteProject(ctx, project)
}
