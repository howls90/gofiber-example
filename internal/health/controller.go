package health

import (
	"github.com/gofiber/fiber/v2"
)

type Controller struct {
	app 		fiber.Router
	service 	*HealthService
}

func (c *Controller) routes() {

	route := c.app.Group("/api/health")

	route.Get("/liveness", func(ctx *fiber.Ctx) error {
		return ctx.JSON(fiber.Map{"message": "ok"})
	})

	route.Get("/readiness", func(ctx *fiber.Ctx) error {
		c.service.Readiness()
		return ctx.JSON(fiber.Map{"message": "ok"})
	})
}