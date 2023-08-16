package pkg

import (
	"context"
	"github.com/Gontafi/golang_jira_analog/pkg/myjira/models"
	q "github.com/Gontafi/golang_jira_analog/pkg/myjira/queries"
	"github.com/jackc/pgx/v5"
	"log/slog"
)

type PriorityRepos struct {
	db  *pgx.Conn
	ctx context.Context
}

func NewPriorityRepos(ctx context.Context, db *pgx.Conn) *PriorityRepos {
	return &PriorityRepos{
		db:  db,
		ctx: ctx,
	}
}

func (r *PriorityRepos) GetById(id int) (models.Priority, error) {
	var priority models.Priority
	err := r.db.QueryRow(r.ctx, q.GetPriorityById, id).Scan(
		&priority.ID, &priority.Name, &priority.Description)
	if err != nil {
		slog.Error("Failed on Priority repository")
		return models.Priority{}, err
	}
	return priority, nil
}

func (r *PriorityRepos) Create(priority models.Priority) (int, error) {
	var id int
	err := r.db.QueryRow(r.ctx, q.CreatePriority,
		priority.Name, priority.Description).Scan(&id)
	if err != nil {
		slog.Error("Failed to create Priority:", err)
		return 0, err
	}
	return id, nil
}

func (r *PriorityRepos) GetAll() ([]models.Priority, error) {
	rows, err := r.db.Query(r.ctx, q.GetPriorities)
	if err != nil {
		slog.Error("Failed on Priority repository")
		return nil, err
	}
	defer rows.Close()

	var priorities []models.Priority
	for rows.Next() {
		var priority models.Priority
		err := rows.Scan(&priority.ID, &priority.Name, &priority.Description)
		if err != nil {
			slog.Error("Failed on Priority repository")
			return nil, err
		}
		priorities = append(priorities, priority)
	}

	if err := rows.Err(); err != nil {
		slog.Error("Failed on Priority repository")
		return nil, err
	}

	return priorities, nil
}

func (r *PriorityRepos) Update(priority models.Priority) error {
	_, err := r.db.Exec(r.ctx, q.UpdatePriority,
		priority.ID, priority.Name, priority.Description)
	if err != nil {
		slog.Error("failed to update Priority")
		return err
	}
	return nil
}

func (r *PriorityRepos) Delete(id int) error {
	_, err := r.db.Exec(r.ctx, q.DeletePriority, id)
	if err != nil {
		slog.Error("failed to delete priority:", err)
		return err
	}
	return nil
}
