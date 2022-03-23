package posts

import (
	"example/fiber/pkg"

	"github.com/gofiber/fiber/v2"
)

// Configure necessaries services for post module
func Init(a *fiber.App) {
	db := pkg.GetDatabaseConn()
	service := PostService{repository: &PostRepository{repo: db}}
	controller := Controller{service: &service}
	Routes(a, controller)
}