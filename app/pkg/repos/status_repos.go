package repos

import (
	"context"
	"github.com/Gontafi/golang_jira_analog/pkg/models"
	q "github.com/Gontafi/golang_jira_analog/pkg/queries"
	"github.com/jackc/pgx/v5"
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

func (r *StatusRepos) GetByStatusID(id int) (models.Status, error) {
	var status models.Status
	err := r.db.QueryRow(r.ctx, q.GetStatusById, id).Scan(
		&status.ID, &status.Name, &status.Description)
	if err != nil {
		return models.Status{}, err
	}
	return status, nil
}

func (r *StatusRepos) CreateStatus(status models.Status) (int, error) {
	var id int
	err := r.db.QueryRow(r.ctx, q.CreateStatus,
		status.Name, status.Description).Scan(&id)
	if err != nil {
		return 0, err
	}
	return id, nil
}

func (r *StatusRepos) GetAllStatuses() ([]models.Status, error) {
	rows, err := r.db.Query(r.ctx, q.GetStatuses)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var statuses []models.Status
	for rows.Next() {
		var status models.Status
		err := rows.Scan(&status.ID, &status.Name, &status.Description)
		if err != nil {
			return nil, err
		}
		statuses = append(statuses, status)
	}

	if err := rows.Err(); err != nil {
		return nil, err
	}

	return statuses, nil
}

func (r *StatusRepos) UpdateStatus(status models.Status) error {
	_, err := r.db.Exec(r.ctx, q.UpdateStatus,
		status.ID, status.Name, status.Description)
	if err != nil {
		return err
	}
	return nil
}

func (r *StatusRepos) DeleteStatus(id int) error {
	_, err := r.db.Exec(r.ctx, q.DeleteStatus, id)
	if err != nil {
		return err
	}
	return nil
}
