package pkg

import (
	"github.com/Gontafi/golang_jira_analog/pkg/myjira/models"
	"github.com/Gontafi/golang_jira_analog/pkg/myjira/repos/pkg"
	"log/slog"
)

type IssueTypeService struct {
	issueTypeRepo *pkg.IssueTypeRepos
}

func NewIssueTypeService(issueTypeRepo *pkg.IssueTypeRepos) *IssueTypeService {
	return &IssueTypeService{issueTypeRepo}
}

func (s *IssueTypeService) Create(issueType models.IssueType) (int, error) {
	var id int
	id, err := s.issueTypeRepo.Create(issueType)
	if err != nil {
		slog.Error("Error in IssueType service while creating", err)
		return 0, err
	}
	return id, nil
}

func (s *IssueTypeService) GetById(id int) (models.IssueType, error) {
	issueType, err := s.issueTypeRepo.GetById(id)

	if err != nil {
		slog.Error("Error in IssueType service while getting", err)
		return models.IssueType{}, err
	}

	return issueType, nil
}

func (s *IssueTypeService) GetAll() ([]models.IssueType, error) {
	issueTypes, err := s.issueTypeRepo.GetAll()
	if err != nil {
		slog.Error("Error in IssueType service while getting", err)
		return []models.IssueType{}, err
	}
	return issueTypes, nil
}

func (s *IssueTypeService) Update(issueType models.IssueType) error {
	err := s.issueTypeRepo.Update(issueType)
	if err != nil {
		slog.Error("Error in IssueType service while updating", err)
		return err
	}
	return nil
}

func (s *IssueTypeService) Delete(id int) error {
	err := s.issueTypeRepo.Delete(id)
	if err != nil {
		slog.Error("Error in IssueType service while deleting", err)
		return err
	}
	return nil
}
