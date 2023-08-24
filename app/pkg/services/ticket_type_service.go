package services

import (
	"github.com/Gontafi/golang_jira_analog/pkg/models"
	"github.com/Gontafi/golang_jira_analog/pkg/repos"
)

type TicketTypeService struct {
	ticketTypeRepo *repos.TicketTypeRepos
}

func NewTicketTypeService(ticketTypeRepos *repos.TicketTypeRepos) *TicketTypeService {
	return &TicketTypeService{ticketTypeRepos}
}

func (s *TicketTypeService) AddTicketType(ticketType models.TicketType) (int, error) {
	var id int
	id, err := s.ticketTypeRepo.CreateTicketType(ticketType)
	if err != nil {
		return 0, err
	}
	return id, nil
}

func (s *TicketTypeService) GetByTicketTypeID(id int) (models.TicketType, error) {
	ticketType, err := s.ticketTypeRepo.GetByTicketTypeID(id)

	if err != nil {
		return models.TicketType{}, err
	}

	return ticketType, nil
}

func (s *TicketTypeService) GetAllTicketTypes() ([]models.TicketType, error) {
	ticketTypes, err := s.ticketTypeRepo.GetAllTicketTypes()
	if err != nil {
		return []models.TicketType{}, err
	}
	return ticketTypes, nil
}

func (s *TicketTypeService) UpdateTicketType(ticketType models.TicketType) error {
	err := s.ticketTypeRepo.UpdateTicketType(ticketType)
	if err != nil {
		return err
	}
	return nil
}

func (s *TicketTypeService) DeleteTicketType(id int) error {
	err := s.ticketTypeRepo.DeleteTicketType(id)
	if err != nil {
		return err
	}
	return nil
}
