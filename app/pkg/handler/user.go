package handler

import (
	"github.com/Gontafi/golang_jira_analog/pkg/models"
	"log"
	"net/http"
	"strconv"

	"github.com/gofiber/fiber/v2"
)

func (h *Handler) CreateUser(c *fiber.Ctx) error {
	var newUser models.User
	if err := c.BodyParser(&newUser); err != nil {
		log.Println("Error parsing request body:", err)
		return ErrorResponse(c, http.StatusBadRequest, "incorrect input")
	}

	userID, err := h.service.User.AddUser(newUser)
	if err != nil {
		return ErrorResponse(c, http.StatusInternalServerError, "Failed to create user")
	}

	return c.Status(http.StatusCreated).JSON(fiber.Map{"id": userID})
}

func (h *Handler) GetUserByID(c *fiber.Ctx) error {
	idParam := c.Params("id")
	id, err := strconv.Atoi(idParam)
	if err != nil {
		return ErrorResponse(c, http.StatusBadRequest, "Invalid user ID")
	}

	user, err := h.service.User.GetByUserID(id)
	if err != nil {
		return ErrorResponse(c, http.StatusInternalServerError, "Failed to retrieve user")
	}

	return c.Status(http.StatusOK).JSON(user)
}

func (h *Handler) GetAllUsers(c *fiber.Ctx) error {
	users, err := h.service.User.GetAllUsers()
	if err != nil {
		return ErrorResponse(c, http.StatusInternalServerError, "Failed to retrieve users")
	}

	return c.Status(http.StatusOK).JSON(users)
}

func (h *Handler) UpdateUser(c *fiber.Ctx) error {
	var updatedUser models.User
	if err := c.BodyParser(&updatedUser); err != nil {
		return ErrorResponse(c, http.StatusBadRequest, "Bad Request")
	}

	err := h.service.User.UpdateUser(updatedUser)
	if err != nil {
		return ErrorResponse(c, http.StatusInternalServerError, "Failed to update user")
	}

	return c.Status(http.StatusOK).JSON(fiber.Map{"message": "User updated successfully"})
}

func (h *Handler) DeleteUser(c *fiber.Ctx) error {
	idParam := c.Params("id")
	id, err := strconv.Atoi(idParam)
	if err != nil {
		return ErrorResponse(c, http.StatusBadRequest, "Bad Request")
	}

	err = h.service.User.DeleteUser(id)
	if err != nil {
		return ErrorResponse(c, http.StatusInternalServerError, "Failed to delete user")
	}

	return c.Status(http.StatusOK).JSON(fiber.Map{"message": "User deleted successfully"})
}
