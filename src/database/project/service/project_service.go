package service

import (
	"ADPwn/database/project/model"
	"ADPwn/database/project/repository"
	"ADPwn/database/utils"
	"context"
	"fmt"
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

//func (s *ProjectService) ProjectByID(ctx context.Context, id string) (model.Project, error) {
//	return s.repo.ProjectByUID(ctx, id)
//}

func (s *ProjectService) SaveSubnet(ctx context.Context, project model.Project, subnet string) (model.Project, error) {
	IPs, err := utils.GenerateIPs(subnet)

	if err != nil {
		return model.Project{}, fmt.Errorf("error in project service while generating hosts from subnetmask: %w", err)
	}

	var hosts []model.Host

	for _, ip := range IPs {
		hosts = append(hosts, *model.NewHost(ip, project.UID, project.Name))
	}

	project.Domains[0].Hosts = append(project.Domains[0].Hosts, hosts...)

	return s.repo.SaveProject(ctx, project)
}

func (s *ProjectService) SaveHost(ctx context.Context, project model.Project, ip string) (model.Project, error) {

	host := *model.NewHost(ip, project.UID, project.Name)
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
