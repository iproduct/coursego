package dao

import (
	"github.com/iproduct/coursego/modules/model"
)

// UserRepo is a repository for users
type UserRepo interface {
	Find(start, count int) ([]model.User, error)
	FindByID(id int) (*model.User, error)
	Create(user *model.User) (*model.User, error)
	Update(user *model.User) (*model.User, error)
	DeleteByID(id int) (*model.User, error)
}
