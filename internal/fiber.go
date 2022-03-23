package internal

import (
	"example/fiber/pkg"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/recover"
)

// Fiber configuration
var FiberAppConfig = fiber.Config{
	ErrorHandler: func(ctx *fiber.Ctx, err error) error {
		code := fiber.StatusInternalServerError
		message := "Internal Server Error"

		if e, ok := err.(*pkg.MyError); ok {
			code = e.Code
			message = e.Message
		} 

		return ctx.Status(code).JSON(fiber.Map{"message": message})
	},
}

// Recovery configuration
var FiberRecovery = recover.New(recover.Config{
	Next:              nil,
	EnableStackTrace:  true,
})