package main

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/go-chi/chi/v5"
	"github.com/iproduct/coursegopro/09-rest/blog"
	"github.com/iproduct/coursegopro/09-rest/container"
	"github.com/iproduct/coursegopro/09-rest/entities"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"log"
	"net/http"
	"time"
)

//var db *gorm.DB

func main() {
	dsn := "root:root@tcp(127.0.0.1:3306)/golang_projects_gorm?charset=utf8mb4&parseTime=True&loc=Local"
	var err error
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal(err)
	}
	db.AutoMigrate(&entities.User{})
	db.AutoMigrate(&blog.Post{}, &blog.Comment{})

	var postsCount int64 = 0
	result := db.Model(blog.Post{}).Count(&postsCount) // SELECT * FROM users;
	if result.Error != nil {
		log.Fatal(result.Error) // returns error
	}
	fmt.Printf("Found %d posts.\n", postsCount)
	if postsCount == 0 {
		post := blog.Post{
			ID:        "",
			CreatedAt: time.Time{},
			Heading:   "New in Go",
			Author:    "Trayan Iliev",
			Content:   "Generics ...",
			Likes:     5,
			Comments:  []blog.Comment{{"GP", "Nice", ""}},
		}
		result = db.Create(&post)

		if result.Error != nil {
			log.Fatal(result.Error) // returns error
		}
		fmt.Printf("New post created with ID: %d -> %+v\nRows afffected: %d\n",
			post.ID, // returns inserted data's primary key
			post,
			result.RowsAffected, // returns inserted records count
		)
		if err != nil {
			log.Fatal(err)
		}
	}

	var postsRepo blog.PostRepository

	//1)postsRepo = container.NewMongoStore(container.MongoOptions{
	//	URI: "mongodb://localhost:27017",
	//})
	//2) postsRepo = container.NewInMemory()
	postsRepo = container.NewMySQLStore(container.MySQLOptions{
		URI: fmt.Sprintf("%s:%s@tcp(127.0.0.1:3306)/golang_projects_gorm?parseTime=true", "root", "root"), //os.Getenv("DB_USER"), os.Getenv("DB_PASS")),
	})
	if postsRepo.Init() != nil {
		log.Fatal(err)
	}

	SetDBMiddleware := func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			timeoutContext, _ := context.WithTimeout(context.Background(), time.Second)
			ctx := context.WithValue(r.Context(), "DB", db.WithContext(timeoutContext))
			ctxWithRepo := context.WithValue(ctx, "postsRepo", postsRepo)
			next.ServeHTTP(w, r.WithContext(ctxWithRepo))
		})
	}

	r := chi.NewRouter()
	r.Use(SetDBMiddleware)

	r.Get("/users", func(w http.ResponseWriter, r *http.Request) {
		db, _ := r.Context().Value("DB").(*gorm.DB)

		var users []entities.User
		db.Find(&users)

		w.Header().Add("Content Type", "application/json")
		data, err := json.MarshalIndent(users, "", "    ")
		if err != nil {
			log.Printf("JSON marshaling failed: %s", err)
		}
		w.Write(data)
	})

	r.Get("/posts", func(w http.ResponseWriter, r *http.Request) {
		//db, _ := r.Context().Value("DB").(*gorm.DB)
		postsRepository, _ := r.Context().Value("postsRepo").(blog.PostRepository)

		//var users []entities.User
		//db.Preload(clause.Associations).Find(&users)

		posts, err := postsRepository.GetAll()
		if err != nil {
			log.Printf("JSON marshaling failed: %s", err)
		}

		w.Header().Add("Content Type", "application/json")
		//w.Header().Add("Location", r.URL.String()+"/"+strconv.Itoa(newID))
		//w.WriteHeader(http.StatusCreated)
		// lots of db operations
		data, err := json.MarshalIndent(posts, "", "    ")
		if err != nil {
			log.Printf("JSON marshaling failed: %s", err)
		}
		w.Write(data)
	})
	http.ListenAndServe(":8080", r)
}
