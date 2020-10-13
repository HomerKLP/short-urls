package configs

import (
	"context"
	"log"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

// Collection обозначает коллекцию в бд
var Collection *mongo.Collection

// ConnectDB подключает базу и проверяет её работоспособность
func ConnectDB(DBUri string) error {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	client, dbError := mongo.Connect(ctx, options.Client().ApplyURI(DBUri))
	if dbError != nil {
		return dbError
	}

	err := client.Ping(context.TODO(), nil)
	if err != nil {
		return err
	}

	Collection = client.Database("short-urls").Collection("short-urls")
	log.Println("Connected to Database...")

	tokenIndex := mongo.IndexModel{
		Keys:    bson.M{"token": 1},
		Options: nil,
	}
	_, err = Collection.Indexes().CreateOne(context.TODO(), tokenIndex)
	if err != nil {
		log.Fatalln("Error while creating index: ", err)
	} else {
		log.Println("Index initialized...")
	}

	return nil
}
