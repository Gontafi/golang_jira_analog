package pkg

import (
	"context"
	"github.com/Gontafi/golang_jira_analog/pkg/myjira/models"
	q "github.com/Gontafi/golang_jira_analog/pkg/myjira/queries"
	"github.com/jackc/pgx/v5"
	"log/slog"
)

type StatusRepos struct {
	db  *pgx.Conn
	ctx context.Context
}

func NewStatusRepos(ctx context.Context, db *pgx.Conn) *StatusRepos {
	return &StatusRepos{
		db:  db,
		ctx: ctx,
	}
}

func (r *StatusRepos) GetById(id int) (models.Status, error) {
	var status models.Status
	err := r.db.QueryRow(r.ctx, q.GetStatusById, id).Scan(
		&status.ID, &status.Name, &status.Description)
	if err != nil {
		slog.Error("Failed on Status repository")
		return models.Status{}, err
	}
	return status, nil
}

func (r *StatusRepos) Create(status models.Status) (int, error) {
	var id int
	err := r.db.QueryRow(r.ctx, q.CreateStatus,
		status.Name, status.Description).Scan(&id)
	if err != nil {
		slog.Error("Failed to create Status:", err)
		return 0, err
	}
	return id, nil
}

func (r *StatusRepos) GetAll() ([]models.Status, error) {
	rows, err := r.db.Query(r.ctx, q.GetStatuses)
	if err != nil {
		slog.Error("Failed on Status repository")
		return nil, err
	}
	defer rows.Close()

	var statuses []models.Status
	for rows.Next() {
		var status models.Status
		err := rows.Scan(&status.ID, &status.Name, &status.Description)
		if err != nil {
			slog.Error("Failed on Status repository")
			return nil, err
		}
		statuses = append(statuses, status)
	}

	if err := rows.Err(); err != nil {
		slog.Error("Failed on Status repository")
		return nil, err
	}

	return statuses, nil
}

func (r *StatusRepos) Update(status models.Status) error {
	_, err := r.db.Exec(r.ctx, q.UpdateStatus,
		status.ID, status.Name, status.Description)
	if err != nil {
		slog.Error("failed to update Status")
		return err
	}
	return nil
}

func (r *StatusRepos) Delete(id int) error {
	_, err := r.db.Exec(r.ctx, q.DeleteStatus, id)
	if err != nil {
		slog.Error("failed to delete status:", err)
		return err
	}
	return nil
}
