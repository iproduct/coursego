package rest

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"github.com/dgrijalva/jwt-go"
	"github.com/go-playground/locales/en"
	ut "github.com/go-playground/universal-translator"
	"github.com/go-playground/validator/v10"
	en_translations "github.com/go-playground/validator/v10/translations/en"
	"github.com/iproduct/coursego/modules/dao"
	"github.com/iproduct/coursego/modules/model"
	"golang.org/x/crypto/bcrypt"
	"time"

	//"golang.org/x/crypto/bcrypt"
	"log"
	"net/http"
	"strconv"
	//"time"

	// bootstrap the mysql driver
	_ "github.com/go-sql-driver/mysql"
	"github.com/gorilla/mux"
)

// App is the top level application
type App struct {
	Router     *mux.Router
	Users      dao.UserRepo
	Validator  *validator.Validate
	Translator ut.Translator
}

// Init method initializes the App
func (a *App) Init(user, password, dbname string) {
	// Create User repository
	a.Users = dao.NewMysql(user, password, dbname)

	// Create and configure validator and translator
	a.Validator = validator.New()
	eng := en.New()
	var uni *ut.UniversalTranslator
	uni = ut.New(eng, eng)
	// this is usually know or extracted from http 'Accept-Language' header
	// also see uni.FindTranslator(...)
	var found bool
	a.Translator, found = uni.GetTranslator("en")
	if !found {
		log.Fatal("translator not found")
	}
	if err := en_translations.RegisterDefaultTranslations(a.Validator, a.Translator); err != nil {
		log.Fatal(err)
	}

	// Create and initialize gorilla/mux router
	a.Router = mux.NewRouter()
	a.initializeRoutes()
}

// Run starts the REST API server
func (a *App) Run(addr string) {
	log.Fatal(http.ListenAndServe(addr, a.Router))
}

func (a *App) initializeRoutes() {
	a.Router.StrictSlash(true)
	//a.Router.Use(CommonMiddleware)
	a.Router.HandleFunc("/users", a.getUsers).Methods("GET")
	a.Router.HandleFunc("/users", a.createUser).Methods("POST")
	a.Router.HandleFunc("/users/{id:[0-9]+}", a.getUserByID).Methods("GET")
	a.Router.HandleFunc("/users/{id:[0-9]+}", a.updateUser).Methods("PUT")
	a.Router.HandleFunc("/users/{id:[0-9]+}", a.deleteUser).Methods("DELETE")
	a.Router.HandleFunc("/login", a.login).Methods("POST")
	// Auth route
	s := a.Router.PathPrefix("/auth").Subrouter()
	s.Use(JwtVerify)
	s.HandleFunc("/users", a.createUser).Methods("POST")
	s.HandleFunc("/users/{id:[0-9]+}", a.getUserByID).Methods("GET")
	s.HandleFunc("/users/{id:[0-9]+}", a.updateUser).Methods("PUT")
	s.HandleFunc("/users/{id:[0-9]+}", a.deleteUser).Methods("DELETE")
}

func (a *App) login(w http.ResponseWriter, r *http.Request) {
	user := &model.User{}
	err := json.NewDecoder(r.Body).Decode(user)
	if err != nil {
		var resp = map[string]interface{}{"status": false, "message": "Invalid request"}
		json.NewEncoder(w).Encode(resp)
		return
	}
	resp, err := a.checkEmailPassword(w, user.Email, user.Password)
	if err == nil {
		json.NewEncoder(w).Encode(resp)
	}
}

func (a *App) checkEmailPassword(w http.ResponseWriter, email, password string) (map[string]interface{}, error)  {

	user, err := a.Users.FindByEmail(email)
	if err != nil {
		respondWithError(w, http.StatusUnauthorized, "Email address not found")
		return nil, err
	}
	expiresAt := time.Now().Add(time.Minute * 100000).Unix()

	err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password))
	if err != nil && err == bcrypt.ErrMismatchedHashAndPassword { //Password does not match!
		respondWithError(w, http.StatusUnauthorized, "Invalid login credentials. Please try again")
		return nil, err
	}

	claims := &model.UserToken{
		UserID: string(user.ID),
		Name:   user.Name,
		Email:  user.Email,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: expiresAt,
			Issuer:    "test",
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	tokenString, error := token.SignedString([]byte("secret"))
	if error != nil {
		fmt.Println(error)
	}

	var resp = map[string]interface{}{"status": false, "message": "logged in"}
	resp["token"] = tokenString //Store the token in the response
	resp["user"] = user
	return resp, nil
}

// User handlers
func (a *App) getUsers(w http.ResponseWriter, r *http.Request) {
	count, _ := strconv.Atoi(r.FormValue("count"))
	start, _ := strconv.Atoi(r.FormValue("start"))

	if count > 10 || count < 1 {
		count = 10
	}
	if start < 0 {
		start = 0
	}

	users, err := a.Users.Find(start, count)
	if err != nil {
		respondWithError(w, http.StatusInternalServerError, err.Error())
		return
	}

	respondWithJSON(w, http.StatusOK, users)
}

func (a *App) createUser(w http.ResponseWriter, r *http.Request) {
	user := &model.User{}

	decoder := json.NewDecoder(r.Body)
	var err error
	if err = decoder.Decode(user); err != nil {
		respondWithError(w, http.StatusBadRequest, "Invalid request payload")
		return
	}
	defer r.Body.Close()

	// Validate User struct
	err = a.Validator.Struct(user)
	if err != nil {
		// translate all error at once
		errs := err.(validator.ValidationErrors)
		respondWithValidationError(errs.Translate(a.Translator), w)
		return
	}

	// Hash the pasword with bcrypt
	pass, err := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost)
	if err != nil {
		fmt.Println(err)
		respondWithError(w, http.StatusInternalServerError, "Password Encryption  failed")
		return
	}
	user.Password = string(pass)

	if user, err = a.Users.Create(user); err != nil {
		respondWithError(w, http.StatusInternalServerError, err.Error())
		return
	}

	respondWithJSON(w, http.StatusCreated, user)
}

func (a *App) getUserByID(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id, err := strconv.Atoi(vars["id"])
	if err != nil {
		respondWithError(w, http.StatusBadRequest, "Invalid user ID")
		return
	}

	var user *model.User
	if user, err = a.Users.FindByID(id); err != nil {
		switch err {
		case sql.ErrNoRows:
			respondWithError(w, http.StatusNotFound, "User not found")
		default:
			respondWithError(w, http.StatusInternalServerError, err.Error())
		}
		return
	}

	respondWithJSON(w, http.StatusOK, user)
}

func (a *App) updateUser(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id, err := strconv.Atoi(vars["id"])
	if err != nil {
		respondWithError(w, http.StatusBadRequest, "Invalid user ID")
		return
	}

	user := &model.User{}
	decoder := json.NewDecoder(r.Body)
	if err := decoder.Decode(user); err != nil {
		respondWithError(w, http.StatusBadRequest, "Invalid resquest payload")
		return
	}
	defer r.Body.Close()
	user.ID = id

	if user, err = a.Users.Update(user); err != nil {
		respondWithError(w, http.StatusInternalServerError, err.Error())
		return
	}

	respondWithJSON(w, http.StatusOK, user)
}

func (a *App) deleteUser(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id, err := strconv.Atoi(vars["id"])
	if err != nil {
		respondWithError(w, http.StatusBadRequest, "Invalid User ID")
		return
	}

	user, err := a.Users.DeleteByID(id)
	if err != nil {
		respondWithError(w, http.StatusInternalServerError, err.Error())
		return
	}

	respondWithJSON(w, http.StatusOK, user)
}

