package main

import (
	"context"
	"encoding/json"
	"github.com/go-chi/chi/v5"
	"github.com/iproduct/coursego/08-databases/entities"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
	"log"
	"net/http"
	"time"
)

var db *gorm.DB

func SetDBMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		timeoutContext, _ := context.WithTimeout(context.Background(), time.Second)
		ctx := context.WithValue(r.Context(), "DB", db.WithContext(timeoutContext))
		next.ServeHTTP(w, r.WithContext(ctx))

	})
}

func main() {
	dsn := "root:root@tcp(127.0.0.1:3306)/golang_projects_gorm?charset=utf8mb4&parseTime=True&loc=Local"
	var err error
	db, err = gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal(err)
	}
	db.AutoMigrate(&entities.User{})

	r := chi.NewRouter()
	r.Use(SetDBMiddleware)

	r.Get("/", func(w http.ResponseWriter, r *http.Request) {
		db, _ := r.Context().Value("DB").(*gorm.DB)

		var users []entities.User
		db.Find(&users)

		// lots of db operations
	})

	r.Get("/users", func(w http.ResponseWriter, r *http.Request) {
		db, _ := r.Context().Value("DB").(*gorm.DB)

		var users []entities.User
		db.Preload(clause.Associations).Find(&users)
		w.Header().Add("Content Type", "application/json")
		//w.Header().Add("Location", r.URL.String()+"/"+strconv.Itoa(newID))
		//w.WriteHeader(http.StatusCreated)
		// lots of db operations
		data, err := json.MarshalIndent(users, "", "    ")
		if err != nil {
			log.Printf("JSON marshaling failed: %s", err)
		}
		w.Write(data)
	})
	http.ListenAndServe(":3100", r)
}
