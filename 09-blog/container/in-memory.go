package container

import (
	"sync"

	"github.com/iproduct/coursego/09-blog/blog"
)

type InMemoryStore struct {
	posts map[string]blog.Post
	mutex sync.RWMutex
}

func NewInMemory() InMemoryStore {
	return InMemoryStore{
		posts: map[string]blog.Post{},
		mutex: sync.RWMutex{},
	}
}

func (c *InMemoryStore) Init() error {
	return nil
}

// GetAll implements blog.Container.
func (c *InMemoryStore) GetAll() ([]blog.Post, error) {
	c.mutex.RLock()
	defer c.mutex.RUnlock()

	posts := []blog.Post{}
	for _, post := range c.posts {
		posts = append(posts, post)
	}
	return posts, nil
}

// Insert implements 09-blog.Container.
func (c *InMemoryStore) Insert(post *blog.Post) error {
	c.mutex.Lock()
	defer c.mutex.Unlock()

	c.posts[post.ID] = *post
	return nil
}

// Delete implements 09-blog.Container.
func (c *InMemoryStore) Delete(id string) error {
	c.mutex.Lock()
	defer c.mutex.Unlock()

	delete(c.posts, id)
	return nil
}
