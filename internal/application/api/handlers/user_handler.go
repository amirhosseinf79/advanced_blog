package handlers

import (
	"github.com/amirhosseinf79/advanced_blog/internal/dto"
	"github.com/amirhosseinf79/advanced_blog/internal/service"
	"github.com/gofiber/fiber/v3"
)

type userHandler struct {
	userService service.UserService
}

func NewUserHandler(userService service.UserService) UserHandler {
	return &userHandler{
		userService: userService,
	}
}

func (h *userHandler) RegisterUser(c fiber.Ctx) error {
	var formData dto.UserRegisterRequest
	response, err := dto.ValidateRequestBody(&formData, c)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(response)
	}

	user, err := h.userService.RegisterUser(formData)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": err.Error(),
		})
	}

	var userResponse dto.UserResponse = dto.UserResponse{
		ID:       user.ID,
		Username: user.Username,
		Email:    user.Email,
	}

	return c.Status(fiber.StatusCreated).JSON(userResponse)
}

func (h *userHandler) LoginUser(c fiber.Ctx) error {
	var formData dto.UserLoginRequest

	response, err := dto.ValidateRequestBody(&formData, c)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(response)
	}

	token, err := h.userService.LoginUser(formData)
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
