package services

import (
	"github.com/Gontafi/golang_jira_analog/pkg/models"
	"github.com/Gontafi/golang_jira_analog/pkg/repos"
)

type StageService struct {
	stageRepo *repos.StageRepos
}

func NewStageService(stageRepo *repos.StageRepos) *StageService {
	return &StageService{stageRepo}
}

func (s *StageService) AddStage(stage models.Stage) (int, error) {
	var id int
	id, err := s.stageRepo.CreateStage(stage)
	if err != nil {
		return 0, err
	}
	return id, nil
}

func (s *StageService) GetByStageID(id int) (models.Stage, error) {
	stage, err := s.stageRepo.GetByStageID(id)

	if err != nil {
		return models.Stage{}, err
	}

	return stage, nil
}

func (s *StageService) GetAllStages() ([]models.Stage, error) {
	stages, err := s.stageRepo.AllStages()
	if err != nil {
		return []models.Stage{}, err
	}
	return stages, nil
}

func (s *StageService) UpdateStage(stage models.Stage) error {
	err := s.stageRepo.UpdateStage(stage)
	if err != nil {
		return err
	}
	return nil
}

func (s *StageService) DeleteStage(id int) error {
	err := s.stageRepo.DeleteStage(id)
	if err != nil {
		return err
	}
	return nil
}
