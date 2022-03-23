package posts

import (
	"github.com/gofiber/fiber/v2"
)


func Routes(a *fiber.App, c Controller) {
	route := a.Group("/api/v1/posts")

	route.Get("", c.getPosts)
	route.Get("/:postId", c.getPost)
	route.Post("", c.createPost)
}