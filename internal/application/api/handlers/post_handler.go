package handlers

import (
	"fmt"

	"github.com/amirhosseinf79/advanced_blog/internal/domain/models"
	"github.com/amirhosseinf79/advanced_blog/internal/dto"
	"github.com/amirhosseinf79/advanced_blog/internal/service"
	"github.com/gofiber/fiber/v3"
)

type postHandler struct {
	postService service.PostService
}

func NewPostHandler(postService service.PostService) PostHandler {
	return &postHandler{postService: postService}
}

func (h *postHandler) CreatePost(c fiber.Ctx) error {
	var formData dto.PostCreateDTO
	var user = c.Locals("user").(models.User)

	response, err := dto.ValidateRequestBody(&formData, c)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(response)
	}
	formData.AuthorID = user.ID

	post, err := h.postService.CreatePost(formData)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": err.Error(),
		})
	}

	var postResponse dto.PostResponse = dto.PostResponse{
		ID:        post.ID,
		Title:     post.Title,
		Content:   post.Content,
		CreatedAt: post.CreatedAt,
		UpdatedAt: post.UpdatedAt,
	}

	return c.Status(fiber.StatusCreated).JSON(postResponse)
}

func (h *postHandler) GetPostByID(c fiber.Ctx) error {
	id := c.Params("id")

	post, err := h.postService.GetPostByID(id)
	if err != nil {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
			"error": "Post not found",
		})
	}

	var postResponse = dto.PostAuthorResponse{
		ID:         post.ID,
		Title:      post.Title,
		Content:    post.Content,
		AuthorName: fmt.Sprintf("%s %s", post.Author.FirstName, post.Author.LastName),
		CreatedAt:  post.CreatedAt,
		UpdatedAt:  post.UpdatedAt,
	}

	return c.Status(fiber.StatusOK).JSON(postResponse)
}

func (h *postHandler) GetAllPosts(c fiber.Ctx) error {
	var filter dto.PostFilterDTO

	response, err := dto.ValidateQueryParams(&filter, c)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(response)
	}

	posts, err := h.postService.GetAllPosts(filter)
	if err != nil || len(posts) == 0 {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
			"error": "Posts not found",
		})
	}

	var postResponses []dto.PostAuthorResponse
	for _, post := range posts {
		postResponse := dto.PostAuthorResponse{
			ID:         post.ID,
			Title:      post.Title,
			Content:    post.Content,
			AuthorName: fmt.Sprintf("%s %s", post.Author.FirstName, post.Author.LastName),
			CreatedAt:  post.CreatedAt,
			UpdatedAt:  post.UpdatedAt,
		}
		postResponses = append(postResponses, postResponse)
	}

	return c.Status(fiber.StatusOK).JSON(postResponses)
}

func (h *postHandler) UpdatePost(c fiber.Ctx) error {
	var formData dto.PostUpdateDTO

	response, err := dto.ValidateRequestBody(&formData, c)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(response)
	}

	post, err := h.postService.UpdatePost(formData)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": err.Error(),
		})
	}

	var postResponse dto.PostResponse = dto.PostResponse{
		ID:        post.ID,
		Title:     post.Title,
		Content:   post.Content,
		CreatedAt: post.CreatedAt,
		UpdatedAt: post.UpdatedAt,
	}

	return c.Status(fiber.StatusOK).JSON(postResponse)
}
