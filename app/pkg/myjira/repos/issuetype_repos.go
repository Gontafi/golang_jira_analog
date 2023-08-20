package repos

import (
	"context"
	"github.com/Gontafi/golang_jira_analog/pkg/myjira/models"
	q "github.com/Gontafi/golang_jira_analog/pkg/myjira/queries"
	"github.com/jackc/pgx/v5"
	"log/slog"
)

type IssueTypeRepos struct {
	db  *pgx.Conn
	ctx context.Context
}

func NewIssueTypeRepos(ctx context.Context, db *pgx.Conn) *IssueTypeRepos {
	return &IssueTypeRepos{
		db:  db,
		ctx: ctx,
	}
}

func (r *IssueTypeRepos) GetByIssueTypeID(id int) (models.IssueType, error) {
	var issueType models.IssueType
	err := r.db.QueryRow(r.ctx, q.GetIssueTypeById, id).Scan(
		&issueType.ID, &issueType.Name, &issueType.Description)
	if err != nil {
		slog.Error("Failed on IssueType repository")
		return models.IssueType{}, err
	}
	return issueType, nil
}

func (r *IssueTypeRepos) CreateIssueType(issueType models.IssueType) (int, error) {
	var id int
	err := r.db.QueryRow(r.ctx, q.CreateIssueType,
		issueType.Name, issueType.Description).Scan(&id)
	if err != nil {
		slog.Error("Failed to create IssueType:", err)
		return 0, err
	}
	return id, nil
}

func (r *IssueTypeRepos) GetAllIssueTypes() ([]models.IssueType, error) {
	rows, err := r.db.Query(r.ctx, q.GetIssueTypes)
	if err != nil {
		slog.Error("Failed on IssueType repository")
		return nil, err
	}
	defer rows.Close()

	var issueTypes []models.IssueType
	for rows.Next() {
		var issueType models.IssueType
		err := rows.Scan(&issueType.ID, &issueType.Name, &issueType.Description)
		if err != nil {
			slog.Error("Failed on IssueType repository")
			return nil, err
		}
		issueTypes = append(issueTypes, issueType)
	}

	if err := rows.Err(); err != nil {
		slog.Error("Failed on IssueType repository")
		return nil, err
	}

	return issueTypes, nil
}

func (r *IssueTypeRepos) UpdateIssueType(issueType models.IssueType) error {
	_, err := r.db.Exec(r.ctx, q.UpdateIssueType,
		issueType.ID, issueType.Name, issueType.Description)
	if err != nil {
		slog.Error("failed to update IssueType")
		return err
	}
	return nil
}

func (r *IssueTypeRepos) DeleteIssueType(id int) error {
	_, err := r.db.Exec(r.ctx, q.DeleteIssueType, id)
	if err != nil {
		slog.Error("failed to delete issue type:", err)
		return err
	}
	return nil
}
