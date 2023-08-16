package pkg

import (
	"github.com/Gontafi/golang_jira_analog/pkg/myjira/models"
	"github.com/Gontafi/golang_jira_analog/pkg/myjira/repos/pkg"
	"log/slog"
)

type PriorityService struct {
	priorityRepo *pkg.PriorityRepos
}

func NewPriorityService(priorityRepo *pkg.PriorityRepos) *PriorityService {
	return &PriorityService{priorityRepo}
}

func (s *PriorityService) Create(priority models.Priority) (int, error) {
	var id int
	id, err := s.priorityRepo.Create(priority)
	if err != nil {
		slog.Error("Error in Priority service while creating", err)
		return 0, err
	}
	return id, nil
}

func (s *PriorityService) GetById(id int) (models.Priority, error) {
	priority, err := s.priorityRepo.GetById(id)

	if err != nil {
		slog.Error("Error in Priority service while getting", err)
		return models.Priority{}, err
	}

	return priority, nil
}

func (s *PriorityService) GetAll() ([]models.Priority, error) {
	priorities, err := s.priorityRepo.GetAll()
	if err != nil {
		slog.Error("Error in Priority service while getting", err)
		return []models.Priority{}, err
	}
	return priorities, nil
}

func (s *PriorityService) Update(priority models.Priority) error {
	err := s.priorityRepo.Update(priority)
	if err != nil {
		slog.Error("Error in Priority service while updating", err)
		return err
	}
	return nil
}

func (s *PriorityService) Delete(id int) error {
	err := s.priorityRepo.Delete(id)
	if err != nil {
		slog.Error("Error in Priority service while deleting", err)
		return err
	}
	return nil
}
