package plugins

import (
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/monitor"
)

// Temetry dashboard configuration
func InitMonitor(a *fiber.App){
	a.Get("/dashboard", monitor.New())
}