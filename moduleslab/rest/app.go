package rest

import (
	"github.com/gorilla/mux"
	"github.com/iproduct/coursego/moduleslab/dao"
	"log"
	"net/http"
)

type App struct {
	Router *mux.Router
	Users  dao.UserRepo
}

func (a *App) Init(user, password, dbname string) {
	a.Users = dao.NewUserRepoMysql(user, password, dbname)
	a.Router = mux.NewRouter()
}

func (a *App) Run(addr string) {
	log.Fatal(http.ListenAndServe(addr, a.Router))
}
