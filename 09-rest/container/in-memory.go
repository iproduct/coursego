package container

import (
	"github.com/iproduct/coursegopro/09-rest/blog"
	"sync"
	"time"
)

type InMemory struct {
	posts map[string]blog.Post
	mutex sync.RWMutex
}

func NewInMemory() *InMemory {
	return &InMemory{
		posts: map[string]blog.Post{},
		mutex: sync.RWMutex{},
	}
}

func (c *InMemory) Init() error {
	err := c.Insert(&blog.Post{
		ID:        "",
		CreatedAt: time.Time{},
		Heading:   "New in Go",
		Author:    "Trayan Iliev",
		Content:   "Generics ...",
		Likes:     5,
		Comments:  []blog.Comment{{"GP", "Nice ..."}},
	})
	return err
}

// GetAll implements blog.Container.
func (c *InMemory) GetAll() ([]blog.Post, error) {
	c.mutex.RLock()
	defer c.mutex.RUnlock()

	posts := []blog.Post{}
	for _, post := range c.posts {
		posts = append(posts, post)
	}
	return posts, nil
}

// Insert implements 09-blog.Container.
func (c *InMemory) Insert(post *blog.Post) error {
	c.mutex.Lock()
	defer c.mutex.Unlock()

	c.posts[post.ID] = *post
	return nil
}

// Delete implements 09-blog.Container.
func (c *InMemory) Delete(id string) error {
	c.mutex.Lock()
	defer c.mutex.Unlock()

	delete(c.posts, id)
	return nil
}
