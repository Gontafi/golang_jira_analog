package queries

const (
	CreateTicket = `INSERT INTO tickets VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9, $10, $11, $12)`

	UpdateTicket = `UPDATE tickets SET project_id=$2, ticket_type_id=$3, 
           ticket_summary=$4, ticket_description=$5, reporter_id=$6, assignee_id=$7,
           status_id=$8, stage_id=$9, priority_id=$10, updated_at=$11, updated_at=$12,
           resolved_at=$13 WHERE id=$1`

	DeleteTicket = `DELETE FROM tickets WHERE ID = $1`

	GetTickets = `SELECT * FROM tickets`

	GetTicketById = `SELECT * FROM tickets WHERE ID = $1`

	// Ticket Type

	CreateTicketType = `INSERT INTO ticket_types VALUES ($1, $2)`

	UpdateTicketType = `UPDATE ticket_types SET name=$2, description=$3 WHERE ID = $1`

	DeleteTicketType = `DELETE FROM ticket_types WHERE ID = $1`

	GetTicketTypes = `SELECT (ID, Name, Description) FROM ticket_types`

	GetTicketTypeById = `SELECT * FROM ticket_types WHERE ID = $1`
)
