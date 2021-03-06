package dao

import (
	"database/sql"
	"errors"
	"github.com/iproduct/coursego/modules/model"
)

type userRepoMock struct {
	users map[int]model.User
	db    *sql.DB
}

//FindAll returns all users
func (r *userRepoMock) Find(start, count int) ([]model.User, error) {
	return nil, errors.New("Not implemented")
}

//FindById return users by user ID or error otherwise
func (r *userRepoMock) FindByID(id int) (*model.User, error) {
	return nil, errors.New("Not implemented")
}

//Create creates and returns new user with autogenerated ID
func (r *userRepoMock) Create(user *model.User) (*model.User, error) {
	return nil, errors.New("Not implemented")
}

//Update updates existing user data
func (r *userRepoMock) Update(user *model.User) (*model.User, error) {
	return nil, errors.New("Not implemented")
}

//DeleteById removes and returns user with specified ID or error otherwise
func (r *userRepoMock) DeleteByID(id int) (*model.User, error) {
	return nil, errors.New("Not implemented")
}

// NewMock is a UserRepo constructor
func NewMock() UserRepo {
	return &userRepoMock{
		users: make(map[int]model.User),
	}
}
