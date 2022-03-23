package posts

import (
	"github.com/google/uuid"

	"example/fiber/pkg"
)

type Post struct {
	Id			uuid.UUID 	`json:"id" gorm:"type:uuid"`
	Title		string		`json:"title" validate:"required"`
	SubTitle	string		`json:"subtitle" validate:"required"`
	Text		string		`json:"text" validate:"required"`
}

type PostRepository struct {
	repo	*pkg.DbRespository
}

func (p *PostRepository) Create(post Post) (*Post, *pkg.MyError) {
	post.Id = uuid.New()

	if err := p.repo.Db.Create(&post).Error; err != nil {
		panic(err)
	}

	return &post, nil
}

func (p *PostRepository) FindById(id string) (*Post, *pkg.MyError) {
	var post Post

	if p.repo.Db.Find(&post, "id = ?", id); post.Id == uuid.Nil {
		return nil, pkg.NotFoundError()  
    }

	return &post, nil
}

func (p *PostRepository) GetAll(offset int, limit int) (*[]Post, *pkg.MyError) {
	var posts []Post

	if p.repo.Db.Offset(offset).Limit(limit).Find(&posts); len(posts) == 0 {
		return nil, pkg.NoContentError()
	}

	return &posts, nil
}