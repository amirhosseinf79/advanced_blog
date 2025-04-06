package handlers

import (
	"github.com/amirhosseinf79/advanced_blog/internal/dto"
	"github.com/amirhosseinf79/advanced_blog/internal/service"
	"github.com/gofiber/fiber/v3"
)

type tokenHandler struct {
	tokenService service.TokenService
}

func NewTokenHandler(tokenService service.TokenService) TokenHandler {
	return &tokenHandler{
		tokenService: tokenService,
	}
}

func (h *tokenHandler) RefreshToken(c fiber.Ctx) error {
	var formData dto.RefreshTokenRequest

	response, err := dto.ValidateRequestBody(&formData, c)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(response)
	}

	token, err := h.tokenService.RefreshToken(formData.RefreshToken)
	if err != nil {
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
			"error": err.Error(),
		})
	}

	var userResponse = dto.UserTokenResponse{
		Token:        token.UUID.String(),
		RefreshToken: token.RefreshToken.String(),
	}

	return c.Status(fiber.StatusOK).JSON(userResponse)
}
