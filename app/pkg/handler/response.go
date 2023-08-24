package handler

import (
	"github.com/gofiber/fiber/v2"
	"log"
)

type Response struct {
	Error string `json:"error"`
}

func ErrorResponse(c *fiber.Ctx, statusCode int, errorMessage string) error {
	err := c.Status(statusCode).JSON(Response{errorMessage})
	if err != nil {
		log.Println("Failed to send response:", err)
		return err
	}
	return nil
}
