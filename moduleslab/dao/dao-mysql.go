package dao

import (
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"github.com/iproduct/coursego/moduleslab/model"
	"log"
)

type userRepoMsql struct {
	db *sql.DB
}

func (u userRepoMsql) Find(start, count int) ([]model.User, error) {
	panic("implement me")
}

func (u userRepoMsql) FindByID(id int) (*model.User, error) {
	panic("implement me")
}

func (u userRepoMsql) FindByEmail(id int) (*model.User, error) {
	panic("implement me")
}

func (u userRepoMsql) Create(user *model.User) (*model.User, error) {
	panic("implement me")
}

func (u userRepoMsql) Update(user *model.User) (*model.User, error) {
	panic("implement me")
}

func (u userRepoMsql) DeleteByID(id int) (*model.User, error) {
	panic("implement me")
}

func NewUserRepoMysql(user, password, dbname string) UserRepo {
	connectionString := fmt.Sprintf("%s:%s@%s", user, password, dbname)
	repo := &userRepoMsql{}
	var err error
	repo.db, err = sql.Open("mysql", connectionString)
	if err != nil {
		log.Fatal(err)
	}
	return repo
}
