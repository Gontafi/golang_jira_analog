package pkg

import (
	"github.com/Gontafi/golang_jira_analog/pkg/myjira/models"
	"github.com/Gontafi/golang_jira_analog/pkg/myjira/repos/pkg"
	"log/slog"
	"time"
)

type IssueService struct {
	issueRepo *pkg.IssueRepos
}

func NewIssueService(issueRepo *pkg.IssueRepos) *IssueService {
	return &IssueService{issueRepo}
}

func (s *IssueService) Create(issue models.Issue) (int, error) {
	var id int

	issue.CreatedAt = time.Now()
	issue.UpdatedAt = time.Now()

	id, err := s.issueRepo.Create(issue)
	if err != nil {
		slog.Error("Error in Issue service while creating", err)
		return 0, err
	}
	return id, nil
}

func (s *IssueService) GetById(id int) (models.Issue, error) {
	issue, err := s.issueRepo.GetById(id)

	if err != nil {
		slog.Error("Error in Issue service while getting", err)
		return models.Issue{}, err
	}

	return issue, nil
}

func (s *IssueService) GetAll() ([]models.Issue, error) {
	issues, err := s.issueRepo.GetAll()
	if err != nil {
		slog.Error("Error in Issue service while getting", err)
		return []models.Issue{}, err
	}
	return issues, nil
}

func (s *IssueService) Update(issue models.Issue) error {
	err := s.issueRepo.Update(issue)
	if err != nil {
		slog.Error("Error in Issue service while updating", err)
		return err
	}
	return nil
}

func (s *IssueService) Delete(id int) error {
	err := s.issueRepo.Delete(id)
	if err != nil {
		slog.Error("Error in Issue service while deleting", err)
		return err
	}
	return nil
}
