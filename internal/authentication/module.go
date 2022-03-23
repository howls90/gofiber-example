package authentication

import (
	"github.com/gofiber/fiber/v2"

	"example/fiber/pkg"
)

// Configure services for Authentication module
func Init(a *fiber.App) {
	jwt := pkg.GetJwtService()
	redis := pkg.GetRedisClient()

	service := AuthService{jwtService: jwt, redisService: &redis}
	controller := Controller{service: &service}
	controller.routes(a)
}