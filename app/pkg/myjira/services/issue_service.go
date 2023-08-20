package services

import (
	"github.com/Gontafi/golang_jira_analog/pkg/myjira/models"
	"github.com/Gontafi/golang_jira_analog/pkg/myjira/repos"
	"log/slog"
	"time"
)

type IssueService struct {
	issueRepo *repos.IssueRepos
}

func NewIssueService(issueRepo *repos.IssueRepos) *IssueService {
	return &IssueService{issueRepo}
}

func (s *IssueService) AddIssue(issue models.Issue) (int, error) {
	var id int

	issue.CreatedAt = time.Now()
	issue.UpdatedAt = time.Now()

	id, err := s.issueRepo.CreateIssue(issue)
	if err != nil {
		slog.Error("Error in Issue service while creating", err)
		return 0, err
	}
	return id, nil
}

func (s *IssueService) GetByIssueID(id int) (models.Issue, error) {
	issue, err := s.issueRepo.GetByIssueID(id)

	if err != nil {
		slog.Error("Error in Issue service while getting", err)
		return models.Issue{}, err
	}

	return issue, nil
}

func (s *IssueService) GetAllIssues() ([]models.Issue, error) {
	issues, err := s.issueRepo.GetAllIssues()
	if err != nil {
		slog.Error("Error in Issue service while getting", err)
		return []models.Issue{}, err
	}
	return issues, nil
}

func (s *IssueService) UpdateIssue(issue models.Issue) error {
	err := s.issueRepo.UpdateIssue(issue)
	if err != nil {
		slog.Error("Error in Issue service while updating", err)
		return err
	}
	return nil
}

func (s *IssueService) DeleteIssue(id int) error {
	err := s.issueRepo.DeleteIssue(id)
	if err != nil {
		slog.Error("Error in Issue service while deleting", err)
		return err
	}
	return nil
}
