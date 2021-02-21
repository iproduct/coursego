package dao

import "github.com/iproduct/coursego/12-modules-rest-jwtauth/model"

type UserRepo interface {
	Find(start, count int) ([]model.User, error)
	FindByID(id int) (*model.User, error)
	FindByEmail(email string) (*model.User, error)
	Create(user *model.User) (*model.User, error)
	Update(user *model.User) (*model.User, error)
	DeleteByID(id int) (*model.User, error)
}
