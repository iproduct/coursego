package dao

// UserRepo is a repository for users
type UserRepo interface {
	FindAll() ([]User, error)
	FindById(id int) (User, error)
	Create(user *User) (User, error)
	Update(user *User) (User, error)
	DeleteById(id int) (User, error)
}

type userRepo struct {
	DB *sql.DB
}

func (r *UserRepo) FindAll() ([]User, error) {
	return errors.New("Not implemented")
}
func (r *UserRepo) FindById(id int) (User, error) {
	return errors.New("Not implemented")
}
func (r *UserRepo) Create(user *User) (User, error) {
	return errors.New("Not implemented")
}
func (r *UserRepo) Update(user *User) (User, error) {
	return errors.New("Not implemented")
}
func (r *UserRepo) DeleteById(id int) (User, error) {
	return errors.New("Not implemented")
}

// New is a UserRepo constructor
func New() *UserRepo {
	return &userRepo{}
}
