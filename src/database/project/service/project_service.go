package service

import (
	"ADPwn/database/project/model"
	"ADPwn/database/project/repository"
	"ADPwn/database/utils"
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

func (s *ProjectService) SaveProject(ctx context.Context, project model.Project) (model.Project, error) {
	return s.repo.SaveProject(ctx, project)
}

//func (s *ProjectService) ProjectByID(ctx context.Context, id string) (model.Project, error) {
//	return s.repo.ProjectByUID(ctx, id)
//}

func (s *ProjectService) SaveSubnet(ctx context.Context, project model.Project, subnet string) (model.Project, error) {
	IPs, _ := utils.GenerateIPs(subnet)
	var hosts []model.Host

	for _, ip := range IPs {
		hosts = append(hosts, *model.NewHost(ip, project.UID))
	}

	project.Domains[0].Hosts = append(project.Domains[0].Hosts, hosts...)

	return s.repo.SaveProject(ctx, project)
}

func (s *ProjectService) SaveHost(ctx context.Context, project model.Project, ip string) (model.Project, error) {

	host := model.Host{IP: ip}
	project.Domains[0].Hosts = append(project.Domains[0].Hosts, host)

	return s.repo.SaveProject(ctx, project)
}

func (s *ProjectService) DeleteProject(ctx context.Context, project model.Project) error {
	return s.repo.DeleteProject(ctx, project)
}
