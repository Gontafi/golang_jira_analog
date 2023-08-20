package repos

import (
	"context"
	"github.com/Gontafi/golang_jira_analog/pkg/myjira/models"
	q "github.com/Gontafi/golang_jira_analog/pkg/myjira/queries"
	"github.com/jackc/pgx/v5"
	"log/slog"
)

type IssueRepos struct {
	db  *pgx.Conn
	ctx context.Context
}

func NewIssueRepos(ctx context.Context, db *pgx.Conn) *IssueRepos {
	return &IssueRepos{
		db:  db,
		ctx: ctx,
	}
}

func (r *IssueRepos) GetByIssueID(id int) (models.Issue, error) {
	var issue models.Issue
	err := r.db.QueryRow(r.ctx, q.GetIssueById, id).Scan(
		&issue.ID, &issue.ProjectID, &issue.IssueTypeID, &issue.IssueSummary, &issue.IssueDescription,
		&issue.ReporterID, &issue.AssigneeID, &issue.StageId, &issue.StatusID, &issue.CreatedAt,
		&issue.UpdatedAt, &issue.ResolverAt)
	if err != nil {
		slog.Error("Failed on Issue repository")
		return models.Issue{}, err
	}
	return issue, nil
}

func (r *IssueRepos) CreateIssue(issue models.Issue) (int, error) {
	var id int
	err := r.db.QueryRow(r.ctx, q.CreateIssue,
		issue.ProjectID, issue.IssueTypeID, issue.IssueSummary, issue.IssueDescription,
		issue.ReporterID, issue.AssigneeID, issue.StageId, issue.StatusID, issue.CreatedAt,
		issue.UpdatedAt, issue.ResolverAt).Scan(&id)
	if err != nil {
		slog.Error("Failed to create Issue:", err)
		return 0, err
	}
	return id, nil
}

func (r *IssueRepos) GetAllIssues() ([]models.Issue, error) {
	rows, err := r.db.Query(r.ctx, q.GetIssues)
	if err != nil {
		slog.Error("Failed on Issue repository")
		return nil, err
	}
	defer rows.Close()

	var issues []models.Issue
	for rows.Next() {
		var issue models.Issue
		err := rows.Scan(
			&issue.ID, &issue.ProjectID, &issue.IssueTypeID, &issue.IssueSummary, &issue.IssueDescription,
			&issue.ReporterID, &issue.AssigneeID, &issue.StageId, &issue.StatusID, &issue.CreatedAt,
			&issue.UpdatedAt, &issue.ResolverAt)
		if err != nil {
			slog.Error("Failed on Issue repository")
			return nil, err
		}
		issues = append(issues, issue)
	}

	if err := rows.Err(); err != nil {
		slog.Error("Failed on Issue repository")
		return nil, err
	}

	return issues, nil
}

func (r *IssueRepos) UpdateIssue(issue models.Issue) error {
	_, err := r.db.Exec(r.ctx, q.UpdateIssue,
		issue.ID, issue.ProjectID, issue.IssueTypeID, issue.IssueSummary, issue.IssueDescription,
		issue.ReporterID, issue.AssigneeID, issue.StageId, issue.StatusID, issue.UpdatedAt, issue.ResolverAt)
	if err != nil {
		slog.Error("failed to update Issue")
		return err
	}
	return nil
}

func (r *IssueRepos) DeleteIssue(id int) error {
	_, err := r.db.Exec(r.ctx, q.DeleteIssue, id)
	if err != nil {
		slog.Error("failed to delete issue:", err)
		return err
	}
	return nil
}
