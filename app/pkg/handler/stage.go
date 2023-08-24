package handler

import (
	"github.com/Gontafi/golang_jira_analog/pkg/models"
	"log"
	"net/http"
	"strconv"

	"github.com/gofiber/fiber/v2"
)

func (h *Handler) CreateStage(c *fiber.Ctx) error {
	var stage models.Stage
	if err := c.BodyParser(&stage); err != nil {
		log.Println("Error parsing request body:", err)
		return ErrorResponse(c, http.StatusBadRequest, "incorrect input")
	}

	stageID, err := h.service.Stage.AddStage(stage)
	if err != nil {
		return ErrorResponse(c, http.StatusInternalServerError, "Failed to create stage")
	}

	return c.Status(http.StatusCreated).JSON(fiber.Map{"id": stageID})
}

func (h *Handler) GetAllStages(c *fiber.Ctx) error {
	stages, err := h.service.Stage.GetAllStages()
	if err != nil {
		log.Println("Failed to fetch stages. Error:", err)
		return ErrorResponse(c, http.StatusInternalServerError, "Failed to fetch stages")
	}

	return c.Status(http.StatusOK).JSON(stages)
}

func (h *Handler) GetStageById(c *fiber.Ctx) error {
	id, err := strconv.Atoi(c.Params("id"))
	if err != nil {
		return ErrorResponse(c, http.StatusBadRequest, "Invalid stage ID")
	}

	stage, err := h.service.Stage.GetByStageID(id)
	if err != nil {
		return ErrorResponse(c, http.StatusNotFound, "Stage not found")
	}

	return c.Status(http.StatusOK).JSON(stage)
}

func (h *Handler) UpdateStage(c *fiber.Ctx) error {
	var stage models.Stage
	if err := c.BodyParser(&stage); err != nil {
		return ErrorResponse(c, http.StatusBadRequest, "incorrect input")
	}

	err := h.service.Stage.UpdateStage(stage)
	if err != nil {
		return ErrorResponse(c, http.StatusInternalServerError, "Failed to update stage")
	}

	return c.Status(http.StatusOK).JSON(fiber.Map{"message": "Stage updated successfully"})
}

func (h *Handler) DeleteStage(c *fiber.Ctx) error {
	id, err := strconv.Atoi(c.Params("id"))
	if err != nil {
		return ErrorResponse(c, http.StatusBadRequest, "Invalid stage ID")
	}

	err = h.service.Stage.DeleteStage(id)
	if err != nil {
		return ErrorResponse(c, http.StatusInternalServerError, "Failed to delete stage")
	}

	return c.Status(http.StatusOK).JSON(fiber.Map{"message": "Stage deleted successfully"})
}
