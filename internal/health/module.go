package health

import (
	"example/fiber/pkg"

	"github.com/gofiber/fiber/v2"
	// "example/fiber/src/plugins"
)

// Configure necessaries services for post module
func Init(a *fiber.App) {
	db := pkg.GetDatabaseConn()
	redis := pkg.GetRedisClient()
	service := HealthService{Db: db, Redis: redis}
	controller := Controller{app: a, service: &service}
	controller.routes()
}