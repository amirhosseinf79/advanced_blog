package middleware

import (
	"github.com/amirhosseinf79/advanced_blog/internal/service"
	"github.com/gofiber/fiber/v3"
)

type tokenMiddleware struct {
	tokenUseCase service.TokenService
}

func NewTokenMiddleware(tokenUseCase service.TokenService) AuthMiddleware {
	return &tokenMiddleware{tokenUseCase: tokenUseCase}
}

func (uc *tokenMiddleware) TokenAuthMiddleware(c fiber.Ctx) error {
	token := c.Get("Authorization")
	if token == "" {
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
			"error": "Missing or invalid token",
		})
	}

	tokenModel, err := uc.tokenUseCase.ValidateToken(token)
	if err != nil {
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
			"error": "Invalid token",
		})
	}

	c.Locals("user", tokenModel.User)
	return c.Next()
}
