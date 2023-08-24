package handler

import (
	"fmt"
	"github.com/Gontafi/golang_jira_analog/pkg/models"
	"github.com/gofiber/fiber/v2"
	"log"
	"net/http"
)

type SignInInput struct {
	Username string `json:"username" binding:"required"`
	Password string `json:"password" binding:"required"`
}
type ForgotPasswordInput struct {
	Username string `json:"username" binding:"required"`
	Password string `json:"password" binding:"required"`
	Code     string `json:"code" binding:"required"`
}

func (h *Handler) SignUp(c *fiber.Ctx) error {
	var input models.User
	if err := c.BodyParser(&input); err != nil {
		log.Println("Error parsing request body:", err)
		return ErrorResponse(c, http.StatusBadRequest, "incorrect input")
	}

	id, err := h.service.Auth.CreateUser(input)

	if err != nil {
		log.Println("Error parsing request body:", err)
		return ErrorResponse(c, http.StatusInternalServerError, "error in server")
	}

	return c.Status(http.StatusOK).JSON(fiber.Map{"id": id})
}

func (h *Handler) SignIn(c *fiber.Ctx) error {
	var input SignInInput

	if err := c.BodyParser(&input); err != nil {
		log.Println("Error parsing request body:", err)
		return ErrorResponse(c, http.StatusBadRequest, "incorrect input")
	}

	token, err := h.service.Auth.GenerateToken(input.Username, input.Password)
	if err != nil {
		log.Println("Error while genration Token:", err)
		return ErrorResponse(c, http.StatusInternalServerError, "Error in server")
	}

	return c.Status(http.StatusOK).JSON(fiber.Map{"token": token})
}

func (h *Handler) ForgotPassword(c *fiber.Ctx) error {
	var code string
	emailForm := struct {
		Email string `json:"email"`
	}{}

	err := c.BodyParser(&emailForm)
	if err != nil {
		return err
	}

	user, err := h.service.User.GetUserByEmail(emailForm.Email)
	if err != nil {
		return ErrorResponse(c, http.StatusNotFound, "User not found")
	}

	code, err = h.service.Auth.ForgotPassword(user.Username)
	if err != nil {
		return ErrorResponse(c, http.StatusInternalServerError, "Failed to generate verification code")
	}

	err = sendVerificationCodeByEmail(emailForm.Email, code)
	if err != nil {
		return ErrorResponse(c, http.StatusInternalServerError, "Failed to send verification code")
	}

	return c.Status(http.StatusOK).JSON(Response{"Verification code sent"})

}

func (h *Handler) VerifyCode(c *fiber.Ctx) error {
	var input ForgotPasswordInput
	if err := c.BodyParser(&input); err != nil {
		log.Println("Error parsing request body:", err)
		return ErrorResponse(c, http.StatusBadRequest, "incorrect input")
	}

	err := h.service.Auth.VerifyResetCode(input.Username, input.Code)

	if err != nil {
		return ErrorResponse(c, http.StatusInternalServerError, "Incorrect code")
	}

	err = h.service.Auth.ChangeUserPassword(input.Username, input.Password)
	if err != nil {
		return ErrorResponse(c, http.StatusInternalServerError, "Failed to change password")
	}

	return c.Status(http.StatusOK).JSON(Response{"Successfully changed"})
}

func sendVerificationCodeByEmail(email string, code string) error {
	//TODO add signals
	fmt.Println(code)
	return nil
}
