package handler

import (
	"github.com/Gontafi/golang_jira_analog/pkg/models"
	"log"
	"net/http"
	"strconv"

	"github.com/gofiber/fiber/v2"
)

func (h *Handler) CreateStatus(c *fiber.Ctx) error {
	var status models.Status
	if err := c.BodyParser(&status); err != nil {
		log.Println("Error parsing request body:", err)
		return ErrorResponse(c, http.StatusBadRequest, "incorrect input")
	}

	statusID, err := h.service.Status.AddStatus(status)
	if err != nil {
		return ErrorResponse(c, http.StatusInternalServerError, "Failed to create status")
	}

	return c.Status(http.StatusCreated).JSON(fiber.Map{"id": statusID})
}

func (h *Handler) GetAllStatuses(c *fiber.Ctx) error {
	statuses, err := h.service.Status.GetAllStatuses()
	if err != nil {
		log.Println("Failed to fetch statuses. Error:", err)
		return ErrorResponse(c, http.StatusInternalServerError, "Failed to fetch statuses")
	}

	return c.Status(http.StatusOK).JSON(statuses)
}

func (h *Handler) GetStatusById(c *fiber.Ctx) error {
	id, err := strconv.Atoi(c.Params("id"))
	if err != nil {
		return ErrorResponse(c, http.StatusBadRequest, "Invalid status ID")
	}

	status, err := h.service.Status.GetByStatusID(id)
	if err != nil {
		return ErrorResponse(c, http.StatusNotFound, "Status not found")
	}

	return c.Status(http.StatusOK).JSON(status)
}

func (h *Handler) UpdateStatus(c *fiber.Ctx) error {
	var status models.Status
	if err := c.BodyParser(&status); err != nil {
		return ErrorResponse(c, http.StatusBadRequest, "incorrect input")
	}

	err := h.service.Status.UpdateStatus(status)
	if err != nil {
		return ErrorResponse(c, http.StatusInternalServerError, "Failed to update status")
	}

	return c.Status(http.StatusOK).JSON(fiber.Map{"message": "Status updated successfully"})
}

func (h *Handler) DeleteStatus(c *fiber.Ctx) error {
	id, err := strconv.Atoi(c.Params("id"))
	if err != nil {
		return ErrorResponse(c, http.StatusBadRequest, "Invalid status ID")
	}

	err = h.service.Status.DeleteStatus(id)
	if err != nil {
		return ErrorResponse(c, http.StatusInternalServerError, "Failed to delete status")
	}

	return c.Status(http.StatusOK).JSON(fiber.Map{"message": "Status deleted successfully"})
}
