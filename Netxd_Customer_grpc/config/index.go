package config

import (
	"context"
	"log"
	"netxd_customer/constants"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"
)

var MongoClient *mongo.Client

func ConnectDatabase() (*mongo.Client, error) {
	ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)
	mongoconnection := options.Client().ApplyURI(constants.ConnectionString)

	MongoClient, err := mongo.Connect(ctx, mongoconnection)

	if err != nil {
		log.Fatal(err.Error())
		return nil, err
	}
	if err := MongoClient.Ping(ctx, readpref.Primary()); err != nil {
		log.Fatal(err.Error())
		return nil, err
	}

	return MongoClient, nil
}

func GetCollection(client *mongo.Client, collectionName string) *mongo.Collection {
	collection := client.Database(constants.DBName).Collection(collectionName)
	return collection
}
