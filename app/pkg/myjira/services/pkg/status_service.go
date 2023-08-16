package pkg

import (
	"github.com/Gontafi/golang_jira_analog/pkg/myjira/models"
	"github.com/Gontafi/golang_jira_analog/pkg/myjira/repos/pkg"
	"log/slog"
)

type StatusService struct {
	statusRepo *pkg.StatusRepos
}

func NewStatusService(statusRepo *pkg.StatusRepos) *StatusService {
	return &StatusService{statusRepo}
}

func (s *StatusService) Create(status models.Status) (int, error) {
	var id int
	id, err := s.statusRepo.Create(status)
	if err != nil {
		slog.Error("Error in Status service while creating", err)
		return 0, err
	}
	return id, nil
}

func (s *StatusService) GetById(id int) (models.Status, error) {
	status, err := s.statusRepo.GetById(id)

	if err != nil {
		slog.Error("Error in Status service while getting", err)
		return models.Status{}, err
	}

	return status, nil
}

func (s *StatusService) GetAll() ([]models.Status, error) {
	statuses, err := s.statusRepo.GetAll()
	if err != nil {
		slog.Error("Error in Status service while getting", err)
		return []models.Status{}, err
	}
	return statuses, nil
}

func (s *StatusService) Update(status models.Status) error {
	err := s.statusRepo.Update(status)
	if err != nil {
		slog.Error("Error in Status service while updating", err)
		return err
	}
	return nil
}

func (s *StatusService) Delete(id int) error {
	err := s.statusRepo.Delete(id)
	if err != nil {
		slog.Error("Error in Status service while deleting", err)
		return err
	}
	return nil
}
