package pkg

import (
	"strings"
)

// Check if token is present
func CheckToken(header []byte, redis RedisClient) (*string, error) {
	payload := string(header[:])

	token := strings.Split(payload, "Bearer ")
	if len(token) == 1 {
		return nil, UnauthorizedError()
	}

	id, err := redis.Get(token[1])
	if err != nil {
		return nil, err
	}
	
	return &id, nil
}