package blogapp

import "github.com/iproduct/coursego/fmi-2023-05-my-blogs/model"

type PostRepository interface {
	GetAll() ([]model.Post, error)
	Insert(*model.Post) error
	Delete(string) error
}

type BlogApp struct {
	posts PostRepository
}

func New(posts PostRepository) *BlogApp {
	return &BlogApp{
		posts: posts,
	}
}

func (b *BlogApp) GetAll() ([]model.Post, error) {
	return b.posts.GetAll()
}

func (b *BlogApp) Add(post *model.Post) error {
	return b.posts.Insert(post)
}

func (b *BlogApp) Delete(id string) error {
	return b.posts.Delete(id)
}
