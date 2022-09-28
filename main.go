package main

import (
	"context"
	"fmt"
	"github.com/gin-gonic/gin"
	"go-crud-mongo/student_business"
	"go-crud-mongo/student_storage"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"
	"log"
)

const (
	HOST     = "35.185.180.53"
	PORT     = "27017"
	USER     = "admin"
	PASSWORD = "2444666668888888"
	DB_NAME  = "student_dev_training"
)

var (
	server      *gin.Engine
	ss          student_storage.StudentStorageService
	sc          student_business.StudentController
	ctx         context.Context
	studC       *mongo.Collection
	mongoClient *mongo.Client
	err         error
)

func init() {
	ctx = context.TODO()

	uriConn := "mongodb://" + USER + ":" + PASSWORD + "@" + HOST + ":" + PORT + "/" + DB_NAME + "?authSource=admin"
	option := options.Client().ApplyURI(uriConn)
	mongoClient, err = mongo.Connect(ctx, option)
	if err != nil {
		log.Fatal("error while connecting with mongo", err)
	}

	err = mongoClient.Ping(ctx, readpref.Primary())
	if err != nil {
		log.Fatal("error while trying to ping mongo", err)
	}

	fmt.Println("mongo connection established")

	studC = mongoClient.Database("student_dev_training").Collection("student")
	index := mongo.IndexModel{
		Keys: bson.M{
			"s_id": 1,
		},
	}

	studC.Indexes().CreateOne(ctx, index)
	ss = student_storage.NewStudentStorage(studC, ctx)
	sc = student_business.New(ss)
	server = gin.Default()
}

func main() {
	//uriConn := "mongodb://" + USER + ":" + PASSWORD + "@" + HOST + ":" + PORT + "/" + DB_NAME + "?authSource=admin"
	//option := options.Client().ApplyURI(uriConn)
	//client, err := mongo.NewClient(option)
	//if err != nil {
	//	log.Fatal(err)
	//}
	//ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)
	//err = client.Connect(ctx)
	//if err != nil {
	//	log.Fatal(err)
	//}
	//defer client.Disconnect(ctx)
	//
	//databases, err := client.ListDatabaseNames(ctx, bson.M{})
	//if err != nil {
	//	log.Fatal(err)
	//}
	//fmt.Println(databases)
	defer mongoClient.Disconnect(ctx)

	basePath := server.Group("/v1")
	sc.RegisterUserRoutes(basePath)

	log.Fatal(server.Run(":5678"))
}
