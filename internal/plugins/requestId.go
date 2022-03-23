package plugins

import (
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/requestid"
	"github.com/google/uuid"
)

func InitRequestId(a *fiber.App) {
	a.Use(requestid.New())
    a.Use(requestid.New(requestid.Config{
        Header: "Test-Service-Header",
        Generator: func() string {
            return uuid.New().String()
        },
    }))
}