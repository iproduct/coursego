package blog

import "time"

type Post struct {
	ID        string
	CreatedAt time.Time
	Heading   string
	Author    string
	Content   string
	Likes     int64
	Comments  []Comment
}

type Comment struct {
	Author  string
	Content string
}

type PostContainer interface {
	Init() error
	GetAll() ([]Post, error)
	Insert(*Post) error
	Delete(string) error
}

type Blog struct {
	posts PostContainer
}

func New(posts PostContainer) *Blog {
	return &Blog{
		posts: posts,
	}
}

func (b *Blog) GetAll() ([]Post, error) {
	return b.posts.GetAll()
}

func (b *Blog) NewPost(post *Post) error {
	return b.posts.Insert(post)
}

func (b *Blog) DeletePost(id string) error {
	return b.posts.Delete(id)
}
