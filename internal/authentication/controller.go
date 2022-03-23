package authentication

import (
	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
)

type Controller struct {
	service			*AuthService
}

func (c *Controller) routes(a *fiber.App) {
	route := a.Group("/auth")

	route.Post("/login", c.login)
}

// Get all Posts godoc
// @Summary Get posts
// @Description get all posts
// @Tags Posts
// @Accept  json
// @Produce  json
// @Router /api/v1/posts [get]
func (c *Controller) login(ctx *fiber.Ctx) error {
	var body LoginRequest
	
	if err := ctx.BodyParser(&body); err != nil {
		return ctx.SendStatus(fiber.StatusInternalServerError)
	}

	validate := validator.New()
	if err := validate.Struct(body); err != nil {
		return err
	}

	token := c.service.CreateToken(body.Email, body.Password)

	return ctx.JSON(fiber.Map{"token": token})
}