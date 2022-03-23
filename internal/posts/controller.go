package posts

import (
	"strconv"

	"github.com/go-playground/validator"
	"github.com/gofiber/fiber/v2"
)

type Controller struct {
	service 	ServiceInterface
}

// Get all Posts godoc
// @Summary Get posts
// @Description get all posts
// @Tags Posts
// @Accept  json
// @Produce  json
// @Success 200 {array} Post
// @Router /api/v1/posts [get]
func (c *Controller) getPosts(ctx *fiber.Ctx) error {
	var err error

	offset, err := strconv.Atoi(ctx.Query("offset", "0"))
	if err != nil || offset < 0 {
		return ctx.Status(400).JSON(fiber.Map{"error": "Query param offset must be integer and higher than 0"})
	}
	limit, err := strconv.Atoi(ctx.Query("limit", "10"))
	if err != nil || limit > 20 {
		return ctx.Status(400).JSON(fiber.Map{"error": "Query param limit must be integer and lower than 20"})
	}

	posts, myErr := c.service.GetPosts(offset, limit)
	if err != nil {
		return ctx.Status(myErr.Code).JSON(fiber.Map{"data": myErr.Message})
	}
	
	return ctx.Status(200).JSON(fiber.Map{"data": posts})
}

// Get Post godoc
// @Summary Show post
// @Description Get post by Id
// @Tags Posts
// @Accept  json
// @Produce  json
// @Success 200 {object} Post
// @Param postId path int true "Post Id"
// @Router /api/v1/posts/{postId} [get]
func (c *Controller) getPost(ctx *fiber.Ctx) error {
	postId := ctx.Params("postId")

	post, err := c.service.GetPost(postId)
	if err != nil {
		return ctx.Status(err.Code).JSON(fiber.Map{"data": err.Message})
	}
	
	return ctx.Status(200).JSON(fiber.Map{"data": post})
}

// Create Post godoc
// @Summary Create Post
// @Description Create post
// @Tags Posts
// @Accept  json
// @Produce  json
// @Param title body string true "Title"
// @Param subtitle body string true "Subtitle"
// @Param text body string true "Text"
// @Success 200 {object} Post
// @Router /api/v1/posts [post]
func (c *Controller) createPost(ctx *fiber.Ctx) error { 
	var body Post
	var err error

	if err = ctx.BodyParser(&body); err != nil {
		return ctx.Status(400).JSON(fiber.Map{"message": err})
	}

	validate := validator.New()
	if err := validate.Struct(body); err != nil {
		return ctx.Status(400).JSON(fiber.Map{"message": err})
	}

	post, myErr := c.service.CreatePost(body)
	if err != nil {
		return ctx.Status(myErr.Code).JSON(fiber.Map{"data": myErr.Message})
	}

	return ctx.Status(201).JSON(fiber.Map{"data": post})
}