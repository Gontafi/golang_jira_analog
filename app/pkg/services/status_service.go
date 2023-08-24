package services

import (
	"github.com/Gontafi/golang_jira_analog/pkg/models"
	"github.com/Gontafi/golang_jira_analog/pkg/repos"
)

type StatusService struct {
	statusRepo *repos.StatusRepos
}

func NewStatusService(statusRepo *repos.StatusRepos) *StatusService {
	return &StatusService{statusRepo}
}

func (s *StatusService) AddStatus(status models.Status) (int, error) {
	var id int
	id, err := s.statusRepo.CreateStatus(status)
	if err != nil {
		return 0, err
	}
	return id, nil
}

func (s *StatusService) GetByStatusID(id int) (models.Status, error) {
	status, err := s.statusRepo.GetByStatusID(id)

	if err != nil {
		return models.Status{}, err
	}

	return status, nil
}

func (s *StatusService) GetAllStatuses() ([]models.Status, error) {
	statuses, err := s.statusRepo.GetAllStatuses()
	if err != nil {
		return []models.Status{}, err
	}
	return statuses, nil
}

func (s *StatusService) UpdateStatus(status models.Status) error {
	err := s.statusRepo.UpdateStatus(status)
	if err != nil {
		return err
	}
	return nil
}

func (s *StatusService) DeleteStatus(id int) error {
	err := s.statusRepo.DeleteStatus(id)
	if err != nil {
		return err
	}
	return nil
}
