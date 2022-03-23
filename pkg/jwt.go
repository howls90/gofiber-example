package pkg

import (
	"time"

	jwtware "github.com/gofiber/jwt/v3"
	"github.com/golang-jwt/jwt/v4"
)

type JwtService struct {
	Secret 			string
	ExpTimeInHours 	time.Duration
}

var service *JwtService

func InitJwt(secret string, exp time.Duration) {
	service = &JwtService{
		Secret: secret,
		ExpTimeInHours: exp,
	}
}

func GetJwtService() *JwtService {
	if service == nil {
		panic("Call InitJwt() first!")
	}
	return service
}

func (j *JwtService) ValidateToken() jwtware.Config {
	return jwtware.Config{
		SigningKey: []byte(j.Secret),
	}
}

func (j *JwtService) CreateToken(email string) string {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"email": email,
		"admin": true,
		"exp": time.Now().Add(j.ExpTimeInHours).Unix(),
	})
	t, err := token.SignedString([]byte(j.Secret))
	if err != nil {
		panic(err)
	}

	return t
}