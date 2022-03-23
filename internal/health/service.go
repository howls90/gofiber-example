package health

import (
	"example/fiber/pkg"
)

type HealthService struct {
	Db		*pkg.DbRespository
	Redis 	pkg.RedisClient
}

func (s *HealthService) Readiness() {
	s.Redis.Ping()
	s.Db.PingDatabase()
}