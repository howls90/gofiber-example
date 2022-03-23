package main

import (
	"flag"
	"fmt"
	"os"
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/pprof"

	_ "example/fiber/docs"
	"example/fiber/internal"
	"example/fiber/internal/authentication"
	"example/fiber/internal/health"
	"example/fiber/internal/plugins"
	"example/fiber/internal/posts"
	"example/fiber/pkg"
)

var (
	app = fiber.New(internal.FiberAppConfig)
	redis pkg.RedisClient
	db *pkg.DbRespository
)

// Read cli input
func inputCli() {
	monitoring := flag.Bool("monitoring", false, "Allow monitor application")
	request_id := flag.Bool("request_id", false, "Allow request ID") 
	debug := flag.Bool("debug", false, "Allow request ID") 
	performance := flag.Bool("performance", false, "Check performance") 

	flag.Parse()

	if *monitoring {
		plugins.InitMonitor(app)
	}

	if *request_id {
		plugins.InitRequestId(app)
	}

	if *debug {
		fmt.Println("Debug active")
	}

	if *performance {
		app.Use(pprof.New())
	}
}

// Configure plugins
func configPlugins() {
	plugins.CheckEnv()
	plugins.InitSwagger(app)

	db = pkg.GetDatabaseConn()
	db.Db.AutoMigrate(&posts.Post{})
	redis = pkg.GetRedisClient()
	service := health.HealthService{Db: db, Redis: redis}
	service.Readiness()
	
	pkg.InitJwt(os.Getenv("JWT_SECRET"), time.Hour * 72)

	app.Use(pkg.InitLog())
	app.Use(internal.FiberRecovery)
}

// @title Fiber Example API
// @version 1.0
// @description This is a sample swagger for Fiber
// @termsOfService http://swagger.io/terms/
// @contact.name API Support
// @contact.email fiber@swagger.io
// @license.name Apache 2.0
// @license.url http://www.apache.org/licenses/LICENSE-2.0.html
// @host localhost:3000
// @BasePath /
func main() {

	inputCli()
	configPlugins()

	authentication.Init(app)

	app.Use("/*", func(c *fiber.Ctx) error {
		header := c.Request().Header.Peek("Authorization")
		id, err := pkg.CheckToken(header, redis)
		if err != nil {
			return c.Status(401).JSON(fiber.Map{"message": "Unauthorized"})
		}
		
		c.Locals("id", id)
		
    	return c.Next()
	})

	posts.Init(app)

	app.Use("/*", func(c *fiber.Ctx) error {
		return c.Status(404).JSON(fiber.Map{"message": "Not Found"})
	})

	app.Listen(":3000")
}
