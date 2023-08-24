package handler

import (
	"net/http"
	"strconv"

	"github.com/gofiber/fiber/v2"
)

func (h *Handler) AddUserToProject(c *fiber.Ctx) error {
	userIdParam := c.Params("user_id")
	projectIdParam := c.Params("project_id")

	userId, err := strconv.Atoi(userIdParam)
	if err != nil {
		return ErrorResponse(c, http.StatusBadRequest, "Invalid user ID")
	}

	projectId, err := strconv.Atoi(projectIdParam)
	if err != nil {
		return ErrorResponse(c, http.StatusBadRequest, "Invalid project ID")
	}

	_, err = h.service.UsersProjects.AddUserToProject(userId, projectId)
	if err != nil {
		return ErrorResponse(c, http.StatusInternalServerError, "Failed to add user to project")
	}

	return c.Status(http.StatusCreated).JSON(fiber.Map{"message": "User added to project successfully"})
}

func (h *Handler) GetUsersFromProject(c *fiber.Ctx) error {
	projectIdParam := c.Params("project_id")

	projectId, err := strconv.Atoi(projectIdParam)
	if err != nil {
		return ErrorResponse(c, http.StatusBadRequest, "Invalid project ID")
	}

	users, err := h.service.UsersProjects.GetUsersFromProject(projectId)
	if err != nil {
		return ErrorResponse(c, http.StatusInternalServerError, "Failed to fetch users from project")
	}

	return c.Status(http.StatusOK).JSON(users)
}

func (h *Handler) GetProjectsFromUser(c *fiber.Ctx) error {
	userIdParam := c.Params("user_id")

	userId, err := strconv.Atoi(userIdParam)
	if err != nil {
		return ErrorResponse(c, http.StatusBadRequest, "Invalid user ID")
	}

	projects, err := h.service.UsersProjects.GetProjectFromUsers(userId)
	if err != nil {
		return ErrorResponse(c, http.StatusInternalServerError, "Failed to fetch projects from user")
	}

	return c.Status(http.StatusOK).JSON(projects)
}

func (h *Handler) DeleteUsersFromProject(c *fiber.Ctx) error {
	userIdParam := c.Params("user_id")
	projectIdParam := c.Params("project_id")

	userId, err := strconv.Atoi(userIdParam)
	if err != nil {
		return ErrorResponse(c, http.StatusBadRequest, "Invalid user ID")
	}

	projectId, err := strconv.Atoi(projectIdParam)
	if err != nil {
		return ErrorResponse(c, http.StatusBadRequest, "Invalid project ID")
	}

	err = h.service.UsersProjects.DeleteUsersFromProject(userId, projectId)
	if err != nil {
		return ErrorResponse(c, http.StatusInternalServerError, "Failed to delete user from project")
	}

	return c.Status(http.StatusOK).JSON(fiber.Map{"message": "User deleted from project successfully"})
}
