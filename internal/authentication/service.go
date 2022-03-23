package authentication

import (
	"example/fiber/pkg"
	"time"
)

type AuthService struct {
	jwtService		*pkg.JwtService
	redisService	*pkg.RedisClient
}

func (a *AuthService) CreateToken(email string, password string) string {
	token := a.jwtService.CreateToken(email)
	a.redisService.Set(token, "dsfdsf", 72 * time.Hour)

	return token
}