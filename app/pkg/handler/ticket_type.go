package handler

import (
	"github.com/Gontafi/golang_jira_analog/pkg/models"
	"log"
	"net/http"
	"strconv"

	"github.com/gofiber/fiber/v2"
)

func (h *Handler) CreateTicketType(c *fiber.Ctx) error {
	var issueType models.TicketType
	if err := c.BodyParser(&issueType); err != nil {
		log.Println("Error parsing request body:", err)
		return ErrorResponse(c, http.StatusBadRequest, "incorrect input")
	}

	issueTypeID, err := h.service.TicketType.AddTicketType(issueType)
	if err != nil {
		return ErrorResponse(c, http.StatusInternalServerError, "Failed to create issue type")
	}

	return c.Status(http.StatusCreated).JSON(fiber.Map{"id": issueTypeID})
}

func (h *Handler) GetAllTicketTypes(c *fiber.Ctx) error {
	issueTypes, err := h.service.TicketType.GetAllTicketTypes()
	if err != nil {
		log.Println("Failed to fetch issue types. Error:", err)
		return ErrorResponse(c, http.StatusInternalServerError, "Failed to fetch issue types")
	}

	return c.Status(http.StatusOK).JSON(issueTypes)
}

func (h *Handler) GetTicketTypeById(c *fiber.Ctx) error {
	id, err := strconv.Atoi(c.Params("id"))
	if err != nil {
		return ErrorResponse(c, http.StatusBadRequest, "Invalid issue type ID")
	}

	issueType, err := h.service.TicketType.GetByTicketTypeID(id)
	if err != nil {
		return ErrorResponse(c, http.StatusNotFound, "Ticket type not found")
	}

	return c.Status(http.StatusOK).JSON(issueType)
}

func (h *Handler) UpdateTicketType(c *fiber.Ctx) error {
	var issueType models.TicketType
	if err := c.BodyParser(&issueType); err != nil {
		return ErrorResponse(c, http.StatusBadRequest, "incorrect input")
	}

	err := h.service.TicketType.UpdateTicketType(issueType)
	if err != nil {
		return ErrorResponse(c, http.StatusInternalServerError, "Failed to update issue type")
	}

	return c.Status(http.StatusOK).JSON(fiber.Map{"message": "Ticket type updated successfully"})
}

func (h *Handler) DeleteTicketType(c *fiber.Ctx) error {
	id, err := strconv.Atoi(c.Params("id"))
	if err != nil {
		return ErrorResponse(c, http.StatusBadRequest, "Invalid issue type ID")
	}

	err = h.service.TicketType.DeleteTicketType(id)
	if err != nil {
		return ErrorResponse(c, http.StatusInternalServerError, "Failed to delete issue type")
	}

	return c.Status(http.StatusOK).JSON(fiber.Map{"message": "Ticket type deleted successfully"})
}
