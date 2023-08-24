package handler

import (
	"github.com/Gontafi/golang_jira_analog/pkg/models"
	"log"
	"net/http"
	"strconv"

	"github.com/gofiber/fiber/v2"
)

func (h *Handler) CreateRole(c *fiber.Ctx) error {
	var role models.Role
	if err := c.BodyParser(&role); err != nil {
		log.Println("Error parsing request body:", err)
		return ErrorResponse(c, http.StatusBadRequest, "incorrect input")
	}

	roleID, err := h.service.Role.AddRole(role)
	if err != nil {
		return ErrorResponse(c, http.StatusInternalServerError, "Failed to create role")
	}

	return c.Status(http.StatusCreated).JSON(fiber.Map{"id": roleID})
}

func (h *Handler) GetAllRoles(c *fiber.Ctx) error {
	roles, err := h.service.Role.GetAllRoles()
	if err != nil {
		log.Println("Failed to fetch roles. Error:", err)
		return ErrorResponse(c, http.StatusInternalServerError, "Failed to fetch roles")
	}

	return c.Status(http.StatusOK).JSON(roles)
}

func (h *Handler) GetRoleById(c *fiber.Ctx) error {
	id, err := strconv.Atoi(c.Params("id"))
	if err != nil {
		return ErrorResponse(c, http.StatusBadRequest, "Invalid role ID")
	}

	role, err := h.service.Role.GetByRoleID(id)
	if err != nil {
		return ErrorResponse(c, http.StatusNotFound, "Role not found")
	}

	return c.Status(http.StatusOK).JSON(role)
}

func (h *Handler) UpdateRole(c *fiber.Ctx) error {
	var role models.Role
	if err := c.BodyParser(&role); err != nil {
		return ErrorResponse(c, http.StatusBadRequest, "incorrect input")
	}

	err := h.service.Role.UpdateRole(role)
	if err != nil {
		return ErrorResponse(c, http.StatusInternalServerError, "Failed to update role")
	}

	return c.Status(http.StatusOK).JSON(fiber.Map{"message": "Role updated successfully"})
}

func (h *Handler) DeleteRole(c *fiber.Ctx) error {
	id, err := strconv.Atoi(c.Params("id"))
	if err != nil {
		return ErrorResponse(c, http.StatusBadRequest, "Invalid role ID")
	}

	err = h.service.Role.DeleteRole(id)
	if err != nil {
		return ErrorResponse(c, http.StatusInternalServerError, "Failed to delete role")
	}

	return c.Status(http.StatusOK).JSON(fiber.Map{"message": "Role deleted successfully"})
}
