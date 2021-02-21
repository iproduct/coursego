package infrastructure

import (
	"context"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"log"
)

var MongoClient *mongo.Client
var Mongodb *mongo.Database

func (e *Environment) InitMongoDB() (client *mongo.Client, db *mongo.Database, err error) {
	clientOptions := options.Client().ApplyURI(e.Databases["mongodb"].Connection)
	MongoClient, err = mongo.Connect(context.TODO(), clientOptions)
	err = MongoClient.Ping(context.TODO(), nil)
	if err != nil {
		return
	}
	Mongodb = MongoClient.Database(e.Databases["mongodb"].Name)
	log.Println("Mongodb Ready!!!")
	return
}
