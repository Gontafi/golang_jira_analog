package services

import (
	"github.com/Gontafi/golang_jira_analog/pkg/myjira/models"
	"github.com/Gontafi/golang_jira_analog/pkg/myjira/repos"
	"log/slog"
)

type IssueTypeService struct {
	issueTypeRepo *repos.IssueTypeRepos
}

func NewIssueTypeService(issueTypeRepo *repos.IssueTypeRepos) *IssueTypeService {
	return &IssueTypeService{issueTypeRepo}
}

func (s *IssueTypeService) AddIssueType(issueType models.IssueType) (int, error) {
	var id int
	id, err := s.issueTypeRepo.CreateIssueType(issueType)
	if err != nil {
		slog.Error("Error in IssueType service while creating", err)
		return 0, err
	}
	return id, nil
}

func (s *IssueTypeService) GetByIssueTypeID(id int) (models.IssueType, error) {
	issueType, err := s.issueTypeRepo.GetByIssueTypeID(id)

	if err != nil {
		slog.Error("Error in IssueType service while getting", err)
		return models.IssueType{}, err
	}

	return issueType, nil
}

func (s *IssueTypeService) GetAllIssueTypes() ([]models.IssueType, error) {
	issueTypes, err := s.issueTypeRepo.GetAllIssueTypes()
	if err != nil {
		slog.Error("Error in IssueType service while getting", err)
		return []models.IssueType{}, err
	}
	return issueTypes, nil
}

func (s *IssueTypeService) UpdateIssueType(issueType models.IssueType) error {
	err := s.issueTypeRepo.UpdateIssueType(issueType)
	if err != nil {
		slog.Error("Error in IssueType service while updating", err)
		return err
	}
	return nil
}

func (s *IssueTypeService) DeleteIssueType(id int) error {
	err := s.issueTypeRepo.DeleteIssueType(id)
	if err != nil {
		slog.Error("Error in IssueType service while deleting", err)
		return err
	}
	return nil
}
