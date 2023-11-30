package inmemory

import (
	"github.com/iproduct/coursego/fmi-2023-05-my-blogs/model"
	"sync"
)

type InMemory struct {
	posts map[string]model.Post
	mutex sync.RWMutex
}

func New() *InMemory {
	return &InMemory{
		posts: make(map[string]model.Post),
		mutex: sync.RWMutex{},
	}
}

func (r *InMemory) GetAll() ([]model.Post, error) {
	r.mutex.RLock()
	defer r.mutex.RUnlock()
	posts := []model.Post{}
	for _, post := range r.posts {
		posts = append(posts, post)
	}
	return posts, nil
}

func (r *InMemory) Insert(post *model.Post) error {
	r.mutex.Lock()
	defer r.mutex.Unlock()
	r.posts[post.ID] = *post
	return nil
}

func (r *InMemory) Delete(id string) error {
	r.mutex.Lock()
	defer r.mutex.Unlock()
	delete(r.posts, id)
	return nil
}
