package repos

import (
	"context"
	"github.com/Gontafi/golang_jira_analog/pkg/models"
	q "github.com/Gontafi/golang_jira_analog/pkg/queries"
	"github.com/jackc/pgx/v5"
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

func (r *PriorityRepos) GetByPriorityID(id int) (models.Priority, error) {
	var priority models.Priority
	err := r.db.QueryRow(r.ctx, q.GetPriorityById, id).Scan(
		&priority.ID, &priority.Name)
	if err != nil {
		return models.Priority{}, err
	}
	return priority, nil
}

func (r *PriorityRepos) CreatePriority(priority models.Priority) (int, error) {
	var id int
	err := r.db.QueryRow(r.ctx, q.CreatePriority,
		priority.Name).Scan(&id)
	if err != nil {
		return 0, err
	}
	return id, nil
}

func (r *PriorityRepos) AllPriorities() ([]models.Priority, error) {
	rows, err := r.db.Query(r.ctx, q.GetPriorities)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var priorities []models.Priority
	for rows.Next() {
		var priority models.Priority
		err := rows.Scan(&priority.ID, &priority.Name)
		if err != nil {
			return nil, err
		}
		priorities = append(priorities, priority)
	}

	if err := rows.Err(); err != nil {
		return nil, err
	}

	return priorities, nil
}

func (r *PriorityRepos) UpdatePriority(priority models.Priority) error {
	_, err := r.db.Exec(r.ctx, q.UpdatePriority,
		priority.ID, priority.Name)
	if err != nil {
		return err
	}
	return nil
}

func (r *PriorityRepos) DeletePriority(id int) error {
	_, err := r.db.Exec(r.ctx, q.DeletePriority, id)
	if err != nil {
		return err
	}
	return nil
}
