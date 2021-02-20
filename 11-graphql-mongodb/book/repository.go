package book

import (
	"context"
	"fmt"
	"github.com/google/uuid"
	"log"

	"github.com/iproduct/coursego/11-graphql-mongodb/infrastructure"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func GetBookByID(ctx context.Context, id string) (result interface{}) {
	var book Book
	data := infrastructure.Mongodb.Collection("booklist").FindOne(ctx, bson.M{"id": id})
	data.Decode(&book)
	return book
}

func GetBookByName(ctx context.Context, name string) (result interface{}) {
	var book Book
	data := infrastructure.Mongodb.Collection("booklist").FindOne(ctx, bson.M{"name": name})
	data.Decode(&book)
	return book
}

func GetBookList(ctx context.Context, limit int) (result interface{}) {
	var book Book
	var books []Book

	option := options.Find().SetLimit(int64(limit))

	cur, err := infrastructure.Mongodb.Collection("booklist").Find(ctx, bson.M{}, option)
	defer cur.Close(ctx)
	if err != nil {
		log.Println(err)
		return nil
	}
	for cur.Next(ctx) {
		cur.Decode(&book)
		books = append(books, book)
	}
	return books
}

func InsertBook(ctx context.Context, book Book) error {
	if book.ID == "" {
		book.ID = uuid.New().String()
	}
	_, err := infrastructure.Mongodb.Collection("booklist").InsertOne(ctx, book)
	return err
}

func UpdateBook(ctx context.Context, book Book) error {
	filter := bson.M{"id": book.ID}
	update := bson.M{"$set": book}
	upsertBool := false
	updateOption := options.UpdateOptions{
		Upsert: &upsertBool,
	}
	updateResult, err := infrastructure.Mongodb.Collection("booklist").UpdateOne(ctx, filter, update, &updateOption)
	if err != nil {
		return err
	}
	if updateResult.MatchedCount != 1 {
		return fmt.Errorf("error updating  book: %v: book not found", book)
	}
	if updateResult.ModifiedCount != 1 {
		return fmt.Errorf("error updating  book: %v: book not updated", book)
	}
	return nil
}

func DeleteBook(ctx context.Context, id string) (Book, error) {
	delResult := infrastructure.Mongodb.Collection("booklist").FindOneAndDelete(ctx, bson.M{"id": id})
	var book Book
	if delResult.Decode(&book); delResult.Err() != nil {
		return book, fmt.Errorf("Error deleting book with ID='%s': %v", id, delResult.Err())
	}
	return book, nil
}
