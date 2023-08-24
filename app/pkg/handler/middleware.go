package handler

import (
	"github.com/gofiber/fiber/v2"
	"net/http"
	"strconv"
	"strings"
)

const (
	authorizationHeader = "Authorization"
	userCtx             = "userId"
)

func (h *Handler) userIdentity(c *fiber.Ctx) error {
	header := c.Get(authorizationHeader)
	if header == "" {
		return ErrorResponse(c, http.StatusUnauthorized, "Authorization header is not set")
	}

	headerParts := strings.Split(header, " ")
	if len(headerParts) != 2 || headerParts[0] != "Bearer" {
		return ErrorResponse(c, http.StatusUnauthorized, "Invalid auth header format")
	}

	userId, err := h.service.Auth.ParseToken(headerParts[1])
	if err != nil {
		return ErrorResponse(c, http.StatusUnauthorized, err.Error())
	}

	c.Locals(userCtx, strconv.Itoa(userId))
	return c.Next()
}
