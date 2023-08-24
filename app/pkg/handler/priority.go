package handler

import (
	"github.com/Gontafi/golang_jira_analog/pkg/models"
	"log"
	"net/http"
	"strconv"

	"github.com/gofiber/fiber/v2"
)

func (h *Handler) CreatePriority(c *fiber.Ctx) error {
	var priority models.Priority
	if err := c.BodyParser(&priority); err != nil {
		log.Println("Error parsing request body:", err)
		return ErrorResponse(c, http.StatusBadRequest, "incorrect input")
	}

	priorityID, err := h.service.Priority.AddPriority(priority)
	if err != nil {
		return ErrorResponse(c, http.StatusInternalServerError, "Failed to create priority")
	}

	return c.Status(http.StatusCreated).JSON(fiber.Map{"id": priorityID})
}

func (h *Handler) GetAllPriorities(c *fiber.Ctx) error {
	priorities, err := h.service.Priority.AllPriorities()
	if err != nil {
		return ErrorResponse(c, http.StatusInternalServerError, "Failed to fetch priorities")
	}

	return c.Status(http.StatusOK).JSON(priorities)
}

func (h *Handler) GetPriorityById(c *fiber.Ctx) error {
	idParam := c.Params("id")
	id, err := strconv.Atoi(idParam)
	if err != nil {
		return ErrorResponse(c, http.StatusBadRequest, "Invalid priority ID")
	}

	priority, err := h.service.Priority.GetByPriorityID(id)
	if err != nil {
		return ErrorResponse(c, http.StatusNotFound, "Priority not found")
	}

	return c.Status(http.StatusOK).JSON(priority)
}

func (h *Handler) UpdatePriority(c *fiber.Ctx) error {
	var priority models.Priority
	if err := c.BodyParser(&priority); err != nil {
		return ErrorResponse(c, http.StatusBadRequest, "Bad Request")
	}

	err := h.service.Priority.UpdatePriority(priority)
	if err != nil {
		return ErrorResponse(c, http.StatusInternalServerError, "Failed to update priority")
	}

	return c.Status(http.StatusOK).JSON(fiber.Map{"message": "Priority updated successfully"})
}

func (h *Handler) DeletePriority(c *fiber.Ctx) error {
	idParam := c.Params("id")
	id, err := strconv.Atoi(idParam)
	if err != nil {
		return ErrorResponse(c, http.StatusBadRequest, "Invalid tag ID")
	}

	err = h.service.Priority.DeletePriority(id)
	if err != nil {
		return ErrorResponse(c, http.StatusInternalServerError, "Failed to delete tag")
	}

	return c.Status(http.StatusOK).JSON(fiber.Map{"message": "Priority deleted successfully"})
}
