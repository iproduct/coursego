package dao

import "github.com/iproduct/coursego/10-modules-rest-jwtauth/model"

type UserRepo interface {
	FindAll(start, count int) ([]model.User, error)
	FindByID(id int) (*model.User, error)
	FindByEmail(email string) (*model.User, error)
	Create(user *model.User) (*model.User, error)
	Update(user *model.User) (*model.User, error)
	DeleteByID(id int) (*model.User, error)
	Count() (int, error)
}
