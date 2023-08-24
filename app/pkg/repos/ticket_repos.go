package repos

import (
	"context"
	"github.com/Gontafi/golang_jira_analog/pkg/models"
	q "github.com/Gontafi/golang_jira_analog/pkg/queries"
	"github.com/jackc/pgx/v5"
)

type TicketRepos struct {
	db  *pgx.Conn
	ctx context.Context
}

func NewTicketRepos(ctx context.Context, db *pgx.Conn) *TicketRepos {
	return &TicketRepos{
		db:  db,
		ctx: ctx,
	}
}

func (r *TicketRepos) GetByTicketID(id int) (models.Ticket, error) {
	var ticket models.Ticket
	err := r.db.QueryRow(r.ctx, q.GetTicketById, id).Scan(
		&ticket.ID, &ticket.ProjectID, &ticket.TicketTypeID, &ticket.TicketSummary, &ticket.TicketDescription,
		&ticket.ReporterID, &ticket.AssigneeID, &ticket.StatusID, &ticket.StatusID, &ticket.PriorityID,
		&ticket.CreatedAt, &ticket.UpdatedAt, &ticket.ResolvedAt)
	if err != nil {
		return models.Ticket{}, err
	}
	return ticket, nil
}

func (r *TicketRepos) CreateTicket(ticket models.Ticket) (int, error) {
	var id int
	err := r.db.QueryRow(r.ctx, q.CreateTicket,
		ticket.ID, ticket.ProjectID, ticket.TicketTypeID, ticket.TicketSummary, ticket.TicketDescription,
		ticket.ReporterID, ticket.AssigneeID, ticket.StatusID, ticket.StatusID, ticket.PriorityID,
		ticket.CreatedAt, ticket.UpdatedAt, ticket.ResolvedAt).Scan(&id)
	if err != nil {
		return 0, err
	}
	return id, nil
}

func (r *TicketRepos) GetAllTickets() ([]models.Ticket, error) {
	rows, err := r.db.Query(r.ctx, q.GetTickets)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var tickets []models.Ticket
	for rows.Next() {
		var ticket models.Ticket
		err := rows.Scan(
			&ticket.ID, &ticket.ProjectID, &ticket.TicketTypeID, &ticket.TicketSummary, &ticket.TicketDescription,
			&ticket.ReporterID, &ticket.AssigneeID, &ticket.StatusID, &ticket.StatusID, &ticket.PriorityID,
			&ticket.CreatedAt, &ticket.UpdatedAt, &ticket.ResolvedAt)
		if err != nil {
			return nil, err
		}
		tickets = append(tickets, ticket)
	}

	if err := rows.Err(); err != nil {
		return nil, err
	}

	return tickets, nil
}

func (r *TicketRepos) UpdateTicket(ticket models.Ticket) error {
	_, err := r.db.Exec(r.ctx, q.UpdateTicket,
		ticket.ID, ticket.ProjectID, ticket.TicketTypeID, ticket.TicketSummary, ticket.TicketDescription,
		ticket.ReporterID, ticket.AssigneeID, ticket.StatusID, ticket.StatusID, ticket.PriorityID,
		ticket.CreatedAt, ticket.UpdatedAt, ticket.ResolvedAt)
	if err != nil {
		return err
	}
	return nil
}

func (r *TicketRepos) DeleteTicket(id int) error {
	_, err := r.db.Exec(r.ctx, q.DeleteTicket, id)
	if err != nil {
		return err
	}
	return nil
}
