package posts

import (
	"net/http/httptest"
	"strings"
	"testing"

	"example/fiber/internal"
	"example/fiber/pkg"

	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
)

type MockPostService struct {}

func (s *MockPostService) GetPosts(offset int, limit int) (*[]Post, *pkg.MyError) {
	return &[]Post {
		{
			Id: uuid.New(),
			Title: "sdfsdfsd",
			SubTitle: "sfsfsdds",
			Text: "ffdgfd",
		},
	}, nil
}

func (s *MockPostService) GetPost(id string) (*Post, *pkg.MyError) {
	if id == "1" {
		return &Post {
			Id: uuid.New(),
			Title: "sdfsdfsd",
			SubTitle: "sfsfsdds",
			Text: "ffdgfd",
		}, nil
	}

	panic(pkg.NotFoundError())
}

func (s *MockPostService) CreatePost(post Post) (*Post, *pkg.MyError) {
	return &Post {
		Id: uuid.New(),
		Title: "sdfsdfsd",
		SubTitle: "sfsfsdds",
		Text: "ffdgfd",
	}, nil
}

var app = fiber.New(internal.FiberAppConfig)
var controller = Controller{service: &MockPostService{}}

func init() {
	app.Use(pkg.InitLog())
	app.Use(internal.FiberRecovery)

	Routes(app, controller)
}

func TestGetAll(t *testing.T) {
	t.Run("Successfully", func(t *testing.T) {
		req := httptest.NewRequest("GET", "/api/v1/posts", nil)

		resp, _ := app.Test(req)

		

		if resp.StatusCode != 200 {
			t.Errorf("GetPost() = %v, want %v", resp.StatusCode, 200)
		}

		// bodyBytes, _ := io.ReadAll(resp.Body)
		// bodyString := string(bodyBytes)
		// expect := `{"data":[{"id":"44dfb876-2d8a-477d-aee7-a2e119d2de4e","title":"sdfsdfsd","subtitle":"sfsfsdds","text":"ffdgfd"}]}`
		// if bodyString != expect {
		// 	t.Errorf("GetPost() = %v, want %v", bodyString, expect)
		// }
	})

	t.Run("bad query params", func(t *testing.T) {
		req := httptest.NewRequest("GET", "/api/v1/posts?offset=-1", nil)

		resp, _ := app.Test(req)

		if resp.StatusCode != 400 {
			t.Errorf("GetPost() = %v, want %v", resp.StatusCode, 400)
		}
	})

	t.Run("bad query params", func(t *testing.T) {
		req := httptest.NewRequest("GET", "/api/v1/posts?limit=50", nil)

		resp, _ := app.Test(req)

		if resp.StatusCode != 400 {
			t.Errorf("GetPost() = %v, want %v", resp.StatusCode, 400)
		}
	})
}

func TestGetById(t *testing.T) {
	t.Run("Successfully", func(t *testing.T) {
		req := httptest.NewRequest("GET", "/api/v1/posts/1", nil)
		resp, _ := app.Test(req)

		if resp.StatusCode != 200 {
			t.Errorf("GetPosts() = %v, want %v", resp.StatusCode, 200)
		}
	})

	t.Run("not found", func(t *testing.T) {

		req := httptest.NewRequest("GET", "/api/v1/posts/1dfg", nil)
		resp, _ := app.Test(req)

		if resp.StatusCode != 404 {
			t.Errorf("GetPosts() = %v, want %v", resp.StatusCode, 404)
		}
	})
}

func TestCreate(t *testing.T) {
	t.Run("Successfully", func(t *testing.T) {
		body := strings.NewReader(`{"title": "sdfdsfsd", "subtitle": "sdsfdsfs", "text": "sfdsfds"}`)
		req := httptest.NewRequest("POST", "/api/v1/posts", body)
		req.Header.Set("Content-Type", "application/json")
		resp, _ := app.Test(req)

		if resp.StatusCode != 201 {
			t.Errorf("GetPosts() = %v, want %v", resp.StatusCode, 201)
		}
	})

	t.Run("Bad params", func(t *testing.T) {
		body := strings.NewReader(`{}`)
		req := httptest.NewRequest("POST", "/api/v1/posts", body)
		req.Header.Set("Content-Type", "application/json")
		resp, _ := app.Test(req)

		if resp.StatusCode != 400 {
			t.Errorf("GetPosts() = %v, want %v", resp.StatusCode, 400)
		}
	})
}