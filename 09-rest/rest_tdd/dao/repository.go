package dao
import (
	"github.com/iproduct/coursegopro/09-rest/rest_tdd/model"
)
type UserRepository interface {
	FindAll() ([]model.User, error)
	FindById(userId int64) ([]model.User, error)
	Create(*model.User) error
	Update(*model.User) error
	DeleteById(userId int64) error
	Count() int64
}
