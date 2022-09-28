package main

import (
	"context"
	"fmt"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"log"
	"time"
)

const (
	HOST     = "35.185.180.53"
	PORT     = "27017"
	USER     = "admin"
	PASSWORD = "2444666668888888"
	DB_NAME  = "student_dev_training"
)

func main() {
	uriConn := "mongodb://" + USER + ":" + PASSWORD + "@" + HOST + ":" + PORT + "/" + DB_NAME + "?authSource=admin"
	option := options.Client().ApplyURI(uriConn)
	client, err := mongo.NewClient(option)
	if err != nil {
		log.Fatal(err)
	}
	ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)
	err = client.Connect(ctx)
	if err != nil {
		log.Fatal(err)
	}
	defer client.Disconnect(ctx)

	databases, err := client.ListDatabaseNames(ctx, bson.M{})
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(databases)
}
