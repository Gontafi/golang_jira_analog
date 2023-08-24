package handler

import (
	"github.com/Gontafi/golang_jira_analog/pkg/models"
	"log"
	"net/http"
	"strconv"

	"github.com/gofiber/fiber/v2"
)

func (h *Handler) CreateTicket(c *fiber.Ctx) error {
	var ticket models.Ticket
	if err := c.BodyParser(&ticket); err != nil {
		log.Println("Error parsing request body:", err)
		return ErrorResponse(c, http.StatusBadRequest, "incorrect input")
	}

	id, err := h.service.Ticket.AddTicket(ticket)
	if err != nil {
		return ErrorResponse(c, http.StatusInternalServerError, "Failed to create ticket")
	}

	return c.Status(http.StatusCreated).JSON(fiber.Map{"id": id})
}

func (h *Handler) GetAllTickets(c *fiber.Ctx) error {
	tickets, err := h.service.Ticket.GetAllTickets()
	if err != nil {
		log.Println("Failed to fetch tickets. Error:", err)
		return ErrorResponse(c, http.StatusInternalServerError, "Failed to fetch tickets")
	}

	return c.Status(http.StatusOK).JSON(fiber.Map{"results": len(tickets), "tickets": tickets})
}

func (h *Handler) GetTicketById(c *fiber.Ctx) error {
	id, err := strconv.Atoi(c.Params("id"))
	if err != nil {
		return ErrorResponse(c, http.StatusBadRequest, "Invalid ticket ID")
	}

	ticket, err := h.service.Ticket.GetByTicketID(id)
	if err != nil {
		return ErrorResponse(c, http.StatusNotFound, "Ticket not found")
	}

	return c.Status(http.StatusOK).JSON(fiber.Map{"ticket": ticket})
}

func (h *Handler) UpdateTicket(c *fiber.Ctx) error {
	var ticket models.Ticket
	if err := c.BodyParser(&ticket); err != nil {
		return ErrorResponse(c, http.StatusBadRequest, "incorrect input")
	}

	err := h.service.Ticket.UpdateTicket(ticket)
	if err != nil {
		return ErrorResponse(c, http.StatusInternalServerError, "Failed to update ticket")
	}

	return c.Status(http.StatusOK).JSON(fiber.Map{"message": "Ticket updated successfully"})
}

func (h *Handler) DeleteTicket(c *fiber.Ctx) error {
	id, err := strconv.Atoi(c.Params("id"))
	if err != nil {
		return ErrorResponse(c, http.StatusBadRequest, "Invalid issue ID")
	}

	err = h.service.Ticket.DeleteTicket(id)
	if err != nil {
		return ErrorResponse(c, http.StatusInternalServerError, "Failed to delete issue")
	}

	return c.Status(http.StatusOK).JSON(fiber.Map{"message": "Ticket deleted successfully"})
}
