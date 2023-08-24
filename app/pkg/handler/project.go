package handler

import (
	"github.com/Gontafi/golang_jira_analog/pkg/models"
	"log"
	"net/http"
	"strconv"

	"github.com/gofiber/fiber/v2"
)

func (h *Handler) CreateProject(c *fiber.Ctx) error {
	var project models.Project
	if err := c.BodyParser(&project); err != nil {
		log.Println("Error parsing request body:", err)
		return ErrorResponse(c, http.StatusBadRequest, "incorrect input")
	}

	projectID, err := h.service.Project.AddProject(project)
	if err != nil {
		return ErrorResponse(c, http.StatusInternalServerError, "Failed to create project")
	}

	return c.Status(http.StatusCreated).JSON(fiber.Map{"id": projectID})
}

func (h *Handler) GetAllProjects(c *fiber.Ctx) error {
	projects, err := h.service.Project.GetAllProjects()
	if err != nil {
		log.Println("Failed to fetch projects. Error:", err)
		return ErrorResponse(c, http.StatusInternalServerError, "Failed to fetch projects")
	}

	return c.Status(http.StatusOK).JSON(projects)
}

func (h *Handler) GetProjectById(c *fiber.Ctx) error {
	id, err := strconv.Atoi(c.Params("id"))
	if err != nil {
		return ErrorResponse(c, http.StatusBadRequest, "Invalid project ID")
	}

	project, err := h.service.Project.GetByProjectID(id)
	if err != nil {
		return ErrorResponse(c, http.StatusNotFound, "Project not found")
	}

	return c.Status(http.StatusOK).JSON(project)
}

func (h *Handler) UpdateProject(c *fiber.Ctx) error {
	var project models.Project
	if err := c.BodyParser(&project); err != nil {
		return ErrorResponse(c, http.StatusBadRequest, "incorrect input")
	}

	err := h.service.Project.UpdateProject(project)
	if err != nil {
		return ErrorResponse(c, http.StatusInternalServerError, "Failed to update project")
	}

	return c.Status(http.StatusOK).JSON(fiber.Map{"message": "Project updated successfully"})
}

func (h *Handler) DeleteProject(c *fiber.Ctx) error {
	id, err := strconv.Atoi(c.Params("id"))
	if err != nil {
		return ErrorResponse(c, http.StatusBadRequest, "Invalid project ID")
	}

	err = h.service.Project.DeleteProject(id)
	if err != nil {
		return ErrorResponse(c, http.StatusInternalServerError, "Failed to delete project")
	}

	return c.Status(http.StatusOK).JSON(fiber.Map{"message": "Project deleted successfully"})
}
