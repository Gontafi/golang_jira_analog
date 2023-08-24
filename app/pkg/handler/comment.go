package handler

import (
	"github.com/Gontafi/golang_jira_analog/pkg/models"
	"github.com/gofiber/fiber/v2"
	"log"
	"net/http"
	"strconv"
)

func (h *Handler) CreateComment(c *fiber.Ctx) error {
	var comment models.Comment

	if err := c.BodyParser(&comment); err != nil {
		log.Println("Error parsing request body:", err)
		return ErrorResponse(c, http.StatusBadRequest, "incorrect input")
	}

	comment.UserID = getUserIDFromContext(c)
	commentID, err := h.service.Comment.AddComment(comment)
	if err != nil {
		return ErrorResponse(c, http.StatusInternalServerError, "Failed to create comment")
	}

	return c.Status(http.StatusCreated).JSON(fiber.Map{"id": commentID})
}

func getUserIDFromContext(c *fiber.Ctx) int {
	userID, err := strconv.Atoi(c.Get(userCtx))
	if err != nil {
		log.Println("Failed to get userID from context")
	}
	return userID
}

func (h *Handler) GetAllComments(c *fiber.Ctx) error {
	comments, err := h.service.Comment.GetAllComments()
	if err != nil {
		log.Println("failed to create comment. Error:", err)
		return ErrorResponse(c, http.StatusInternalServerError, "Failed to fetch comments")
	}

	return c.Status(http.StatusOK).JSON(fiber.Map{"results": len(comments), "comments": comments})
}

func (h *Handler) GetCommentById(c *fiber.Ctx) error {
	if c.Params("id") == "" {
		return ErrorResponse(c, http.StatusBadRequest, "No ID given")
	}
	id, err := strconv.Atoi(c.Params("id"))
	if err != nil {
		return ErrorResponse(c, http.StatusBadRequest, "Invalid comment ID")
	}

	comment, err := h.service.Comment.GetByCommentID(id)
	if err != nil {
		return ErrorResponse(c, http.StatusBadRequest, "Comment not found")
	}

	return c.Status(http.StatusOK).JSON(fiber.Map{"comment": comment})
}

func (h *Handler) UpdateComment(c *fiber.Ctx) error {
	var comment models.Comment
	if err := c.BodyParser(&comment); err != nil {
		return ErrorResponse(c, http.StatusBadRequest, "incorrect input")
	}

	err := h.service.Comment.UpdateComment(comment)
	if err != nil {
		return ErrorResponse(c, http.StatusInternalServerError, "Failed to update comment")
	}

	return c.Status(http.StatusOK).JSON(fiber.Map{"message": "Comment updated successfully"})
}

func (h *Handler) DeleteComment(c *fiber.Ctx) error {
	if c.Params("id") == "" {
		return ErrorResponse(c, http.StatusBadRequest, "No ID given")
	}
	id, err := strconv.Atoi(c.Params("id"))
	if err != nil {
		return ErrorResponse(c, http.StatusBadRequest, "Invalid comment ID")
	}

	err = h.service.Comment.DeleteComment(id)
	if err != nil {
		return ErrorResponse(c, http.StatusInternalServerError, "Failed to delete comment")
	}

	return c.Status(http.StatusOK).JSON(fiber.Map{"message": "Comment deleted successfully"})
}
