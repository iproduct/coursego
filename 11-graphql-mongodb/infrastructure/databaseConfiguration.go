package infrastructure

import (
	"context"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"log"
)

var Mongodb *mongo.Database

func (e *Environment) InitMongoDB() (db *mongo.Database, err error) {
	clientOptions := options.Client().ApplyURI(e.Databases["mongodb"].Connection)
	client, err := mongo.Connect(context.TODO(), clientOptions)
	err = client.Ping(context.TODO(), nil)
	if err != nil {
		return db, err
	}
	Mongodb = client.Database(e.Databases["mongodb"].Name)
	log.Println("Mongodb Ready!!!")
	return db, err
}
