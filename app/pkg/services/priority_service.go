package services

import (
	"github.com/Gontafi/golang_jira_analog/pkg/models"
	"github.com/Gontafi/golang_jira_analog/pkg/repos"
)

type PriorityService struct {
	PriorityRepo *repos.PriorityRepos
}

func NewPriorityService(PriorityRepo *repos.PriorityRepos) *PriorityService {
	return &PriorityService{PriorityRepo}
}

func (s *PriorityService) AddPriority(priority models.Priority) (int, error) {
	var id int
	id, err := s.PriorityRepo.CreatePriority(priority)
	if err != nil {
		return 0, err
	}
	return id, nil
}

func (s *PriorityService) GetByPriorityID(id int) (models.Priority, error) {
	priority, err := s.PriorityRepo.GetByPriorityID(id)

	if err != nil {
		return models.Priority{}, err
	}

	return priority, nil
}

func (s *PriorityService) AllPriorities() ([]models.Priority, error) {
	priorities, err := s.PriorityRepo.AllPriorities()
	if err != nil {
		return []models.Priority{}, err
	}
	return priorities, nil
}

func (s *PriorityService) UpdatePriority(priority models.Priority) error {
	err := s.PriorityRepo.UpdatePriority(priority)
	if err != nil {
		return err
	}
	return nil
}

func (s *PriorityService) DeletePriority(id int) error {
	err := s.PriorityRepo.DeletePriority(id)
	if err != nil {
		return err
	}
	return nil
}
