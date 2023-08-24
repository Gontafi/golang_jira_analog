package repos

import (
	"context"
	"github.com/Gontafi/golang_jira_analog/pkg/models"
	q "github.com/Gontafi/golang_jira_analog/pkg/queries"
	"github.com/jackc/pgx/v5"
)

type TicketTypeRepos struct {
	db  *pgx.Conn
	ctx context.Context
}

func NewTicketTypeRepos(ctx context.Context, db *pgx.Conn) *TicketTypeRepos {
	return &TicketTypeRepos{
		db:  db,
		ctx: ctx,
	}
}

func (r *TicketTypeRepos) GetByTicketTypeID(id int) (models.TicketType, error) {
	var ticketType models.TicketType
	err := r.db.QueryRow(r.ctx, q.GetTicketTypeById, id).Scan(
		&ticketType.ID, &ticketType.Name, &ticketType.Description)
	if err != nil {
		return models.TicketType{}, err
	}
	return ticketType, nil
}

func (r *TicketTypeRepos) CreateTicketType(ticketType models.TicketType) (int, error) {
	var id int
	err := r.db.QueryRow(r.ctx, q.CreateTicketType,
		ticketType.Name, ticketType.Description).Scan(&id)
	if err != nil {
		return 0, err
	}
	return id, nil
}

func (r *TicketTypeRepos) GetAllTicketTypes() ([]models.TicketType, error) {
	rows, err := r.db.Query(r.ctx, q.GetTicketTypes)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var ticketTypes []models.TicketType
	for rows.Next() {
		var ticketType models.TicketType
		err := rows.Scan(&ticketType.ID, &ticketType.Name, &ticketType.Description)
		if err != nil {
			return nil, err
		}
		ticketTypes = append(ticketTypes, ticketType)
	}

	if err := rows.Err(); err != nil {
		return nil, err
	}

	return ticketTypes, nil
}

func (r *TicketTypeRepos) UpdateTicketType(ticketType models.TicketType) error {
	_, err := r.db.Exec(r.ctx, q.UpdateTicketType,
		ticketType.ID, ticketType.Name, ticketType.Description)
	if err != nil {
		return err
	}
	return nil
}

func (r *TicketTypeRepos) DeleteTicketType(id int) error {
	_, err := r.db.Exec(r.ctx, q.DeleteTicketType, id)
	if err != nil {
		return err
	}
	return nil
}
