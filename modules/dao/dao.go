package dao

import (
	"github.com/iproduct/coursego/modules/model"
)

// UserRepo is a repository for users
type UserRepo interface {
	FindAll() ([]model.User, error)
	FindById(id int) (model.User, error)
	Create(user *model.User) (model.User, error)
	Update(user *model.User) (model.User, error)
	DeleteById(id int) (model.User, error)
}
