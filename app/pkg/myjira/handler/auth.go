package handler

import (
	"fmt"
	"github.com/Gontafi/golang_jira_analog/pkg/myjira/models"
	"github.com/gin-gonic/gin"
	"net/http"
)

func (h *Handler) SignUp(c *gin.Context) {
	var input models.User
	if err := c.ShouldBindJSON(&input); err != nil {
		NewErrorResponse(c, http.StatusBadRequest, "incorrect input")
		return
	}
	fmt.Println(input.Username)
	id, err := h.service.Auth.CreateUser(input)
	if err != nil {
		NewErrorResponse(c, http.StatusInternalServerError, "Error in server")
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"id": id,
	})
}

type SignInInput struct {
	Username string `json:"username" binding:"required"`
	Password string `json:"password" binding:"required"`
}
type ForgotPasswordInput struct {
	Username string `json:"username" binding:"required"`
	Password string `json:"password" binding:"required"`
	Code     string `json:"code" binding:"required"`
}

func (h *Handler) SignIn(c *gin.Context) {
	var input SignInInput
	if err := c.BindJSON(&input); err != nil {
		NewErrorResponse(c, http.StatusBadRequest, "incorrect input")
		return
	}
	fmt.Println("processing token")

	token, err := h.service.Auth.GenerateToken(input.Username, input.Password)
	if err != nil {
		NewErrorResponse(c, http.StatusInternalServerError, "Error in server")
		return
	}
	fmt.Println("finished, token:", token)

	c.JSON(http.StatusOK, gin.H{
		"token": token,
	})
}

func (h *Handler) ForgotPassword(c *gin.Context) {
	var code string
	email := c.PostForm("email")
	user, err := h.service.User.GetUserByEmail(email)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "User not found"})
		return
	}

	code, err = h.service.Auth.ForgotPassword(user.Username)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to generate verification code"})
		return
	}

	err = sendVerificationCodeByEmail(email, code)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to send verification code"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Verification code sent"})
}

func (h *Handler) VerifyCode(c *gin.Context) {
	var input ForgotPasswordInput
	if err := c.BindJSON(&input); err != nil {
		NewErrorResponse(c, http.StatusBadRequest, "incorrect input")
		return
	}
	err := h.service.Auth.VerifyResetCode(input.Username, input.Code)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to verify code"})
		return
	}

	err = h.service.Auth.ChangeUserPassword(input.Username, input.Password)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to change password"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Successfully changed"})
}
func sendVerificationCodeByEmail(email, code string) error {
	//TODO
	fmt.Println(code)
	return nil
}
