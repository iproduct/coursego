package rest

import (
	"github.com/iproduct/coursego/modules/repository"
	"database/sql"
	// import the mysql driver
	_ "github.com/go-sql-driver/mysql"
	"github.com/gorilla/mux"
)

// App is the top level application
type App struct {
	Router *mux.Router
	
	Users  repository.UserRepo
}

// Init method initializes the App
func (a *App) Init(user, password, dbname string) {}

// Run starts the REST API server
func (a *App) Run(addres string) {}
