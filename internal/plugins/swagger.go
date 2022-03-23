package plugins

import (
	_ "example/fiber/docs"

	swagger "github.com/arsmn/fiber-swagger/v2"
	"github.com/gofiber/fiber/v2"
)

// Configure Swagger namespace
func InitSwagger(a *fiber.App){
	a.Get("/swagger/*", swagger.HandlerDefault)
}