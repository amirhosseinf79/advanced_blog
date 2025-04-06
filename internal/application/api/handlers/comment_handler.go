package handlers

import (
	"github.com/amirhosseinf79/advanced_blog/internal/domain/models"
	"github.com/amirhosseinf79/advanced_blog/internal/dto"
	"github.com/amirhosseinf79/advanced_blog/internal/service"
	"github.com/amirhosseinf79/advanced_blog/internal/shared"
	"github.com/gofiber/fiber/v3"
)

type commentHnadler struct {
	commentService service.CommentService
}

func NewCommentHandler(commentService service.CommentService) CommentHandler {
	return &commentHnadler{commentService: commentService}
}

func (h *commentHnadler) GetComments(c fiber.Ctx) error {
	var filter dto.CommentFilter
	response, err := dto.ValidateQueryParams(&filter, c)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(response)
	}

	comments, total, err := h.commentService.AllPostComments(filter)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": err.Error(),
		})
	}

	if len(comments) == 0 {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
			"message": "No comments found",
		})
	}

	var commentList []dto.CommentResponse

	for _, coment := range comments {
		data := dto.CommentResponse{
			ID:        coment.ID,
			Comment:   coment.Comment,
			Author:    coment.User.Username,
			CreatedAt: coment.CreatedAt,
			UpdatedAt: coment.UpdatedAt,
		}
		commentList = append(commentList, data)
	}

	paginator := shared.NewPaginator(total, filter.Page, filter.PageSize, commentList)
	return c.JSON(paginator.Paginate())
}

func (h *commentHnadler) AddComment(c fiber.Ctx) error {
	var comment dto.CommentCreateDTO
	response, err := dto.ValidateRequestBody(&comment, c)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(response)
	}

	comment.UserID = c.Locals("user").(models.User).ID
	commentModel, err := h.commentService.AddComment(comment)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": err.Error(),
		})
	}
	var commentResponse = dto.CommentResponse{
		ID:        commentModel.ID,
		Comment:   commentModel.Comment,
		CreatedAt: commentModel.CreatedAt,
		UpdatedAt: commentModel.UpdatedAt,
	}
	return c.JSON(commentResponse)
}
