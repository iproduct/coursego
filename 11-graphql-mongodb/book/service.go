package book

import (
	"context"
	"fmt"
	"log"
	"net/http"

	"github.com/iproduct/coursego/11-graphql-mongodb/infrastructure"

	"github.com/go-chi/chi"
	"go.mongodb.org/mongo-driver/bson"
)

/* Rest API */
func RestApiGetBookAllBooks(w http.ResponseWriter, r *http.Request) {
	ctx := context.TODO()
	books := []Book{}
	//var results []bson.M

	cur, err := infrastructure.Mongodb.Collection("booklist").Find(ctx, bson.M{})
	defer cur.Close(ctx)
	if err != nil {
		log.Println(err)
		return
	}
	if err = cur.All(context.TODO(), &books); err != nil {
		log.Fatal(err)
	}
	for _, result := range books {
		fmt.Println(result)
	}

	if s := books; s != nil {
		HttpResponseSuccess(w, r, books)
		return
	}
	return
}

func RestApiGetBook(w http.ResponseWriter, r *http.Request) {
	var ctx context.Context
	books := &Book{}
	bookName := chi.URLParam(r, "bookname")

	cur, err := infrastructure.Mongodb.Collection("booklist").Find(ctx, bson.M{"name": bookName})
	defer cur.Close(ctx)
	if err != nil {
		log.Println(err)
		return
	}
	for cur.Next(ctx) {
		cur.Decode(&books)
	}

	if s := books; s != nil {
		HttpResponseSuccess(w, r, books)
		return
	}
	return
}
