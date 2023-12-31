package services

import (
	"github.com/Gontafi/golang_jira_analog/pkg/models"
	"github.com/Gontafi/golang_jira_analog/pkg/repos"
)

type ProjectService struct {
	projectRepo *repos.ProjectRepos
}

func NewProjectService(projectRepo *repos.ProjectRepos) *ProjectService {
	return &ProjectService{projectRepo}
}

func (s *ProjectService) AddProject(project models.Project) (int, error) {
	var id int

	id, err := s.projectRepo.CreateProject(project)
	if err != nil {
		return 0, err
	}
	return id, nil
}

func (s *ProjectService) GetByProjectID(id int) (models.Project, error) {
	project, err := s.projectRepo.GetByProjectID(id)

	if err != nil {
		return models.Project{}, err
	}

	return project, nil
}

func (s *ProjectService) GetAllProjects() ([]models.Project, error) {
	projects, err := s.projectRepo.GetAllProjects()
	if err != nil {
		return []models.Project{}, err
	}
	return projects, nil
}

func (s *ProjectService) UpdateProject(project models.Project) error {
	err := s.projectRepo.UpdateProject(project)
	if err != nil {
		return err
	}
	return nil
}

func (s *ProjectService) DeleteProject(id int) error {
	err := s.projectRepo.DeleteProject(id)
	if err != nil {
		return err
	}
	return nil
}
