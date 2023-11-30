package inmemory

import (
	"github.com/iproduct/coursego/fmi-2023-05-my-blogs/model"
	"sync"
)

type InMemoryRepository struct {
	posts map[string]model.Post
	mutex sync.RWMutex
}

func New() *InMemoryRepository {
	return &InMemoryRepository{
		posts: make(map[string]model.Post),
		mutex: sync.RWMutex{},
	}
}

func (r *InMemoryRepository) GetAll() ([]model.Post, error) {
	r.mutex.RLock()
	defer r.mutex.RUnlock()
	posts := []model.Post{}
	for _, post := range r.posts {
		posts = append(posts, post)
	}
	return posts, nil
}

func (r *InMemoryRepository) Insert(post *model.Post) error {
	r.mutex.Lock()
	defer r.mutex.Unlock()
	r.posts[post.ID] = *post
	return nil
}

func (r *InMemoryRepository) Delete(id string) error {
	r.mutex.Lock()
	defer r.mutex.Unlock()
	delete(r.posts, id)
	return nil
}
