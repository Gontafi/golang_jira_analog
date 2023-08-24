package services

import (
	"github.com/Gontafi/golang_jira_analog/pkg/models"
	"github.com/Gontafi/golang_jira_analog/pkg/repos"
	"time"
)

type TicketService struct {
	ticketRepos *repos.TicketRepos
}

func NewTicketService(ticketRepos *repos.TicketRepos) *TicketService {
	return &TicketService{ticketRepos}
}

func (s *TicketService) AddTicket(ticket models.Ticket) (int, error) {
	var id int

	ticket.CreatedAt = time.Now()
	ticket.UpdatedAt = time.Now()

	id, err := s.ticketRepos.CreateTicket(ticket)
	if err != nil {
		return 0, err
	}
	return id, nil
}

func (s *TicketService) GetByTicketID(id int) (models.Ticket, error) {
	ticket, err := s.ticketRepos.GetByTicketID(id)

	if err != nil {
		return models.Ticket{}, err
	}

	return ticket, nil
}

func (s *TicketService) GetAllTickets() ([]models.Ticket, error) {
	tickets, err := s.ticketRepos.GetAllTickets()
	if err != nil {
		return []models.Ticket{}, err
	}
	return tickets, nil
}

func (s *TicketService) UpdateTicket(ticket models.Ticket) error {
	err := s.ticketRepos.UpdateTicket(ticket)
	if err != nil {
		return err
	}
	return nil
}

func (s *TicketService) DeleteTicket(id int) error {
	err := s.ticketRepos.DeleteTicket(id)
	if err != nil {
		return err
	}
	return nil
}
