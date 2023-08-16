package pkg

import (
	"github.com/Gontafi/golang_jira_analog/pkg/myjira/models"
	"github.com/Gontafi/golang_jira_analog/pkg/myjira/repos/pkg"
	"log/slog"
)

type ProjectService struct {
	projectRepo *pkg.ProjectRepos
}

func NewProjectService(projectRepo *pkg.ProjectRepos) *ProjectService {
	return &ProjectService{projectRepo}
}

func (s *ProjectService) Create(project models.Project) (int, error) {
	var id int

	id, err := s.projectRepo.Create(project)
	if err != nil {
		slog.Error("Error in Project service while creating", err)
		return 0, err
	}
	return id, nil
}

func (s *ProjectService) GetById(id int) (models.Project, error) {
	project, err := s.projectRepo.GetById(id)

	if err != nil {
		slog.Error("Error in Project service while getting", err)
		return models.Project{}, err
	}

	return project, nil
}

func (s *ProjectService) GetAll() ([]models.Project, error) {
	projects, err := s.projectRepo.GetAll()
	if err != nil {
		slog.Error("Error in Project service while getting", err)
		return []models.Project{}, err
	}
	return projects, nil
}

func (s *ProjectService) Update(project models.Project) error {
	err := s.projectRepo.Update(project)
	if err != nil {
		slog.Error("Error in Project service while updating", err)
		return err
	}
	return nil
}

func (s *ProjectService) Delete(id int) error {
	err := s.projectRepo.Delete(id)
	if err != nil {
		slog.Error("Error in Project service while deleting", err)
		return err
	}
	return nil
}
