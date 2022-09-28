package main

import (
	"context"
	"fmt"
	"github.com/gin-gonic/gin"
	"go-crud-mongo/student_business"
	"go-crud-mongo/student_model"
	"go-crud-mongo/student_storage"
	"go-crud-mongo/utils"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"
	"log"
	"strconv"
	"sync"
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

var wg sync.WaitGroup
var mu sync.Mutex

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
		Options: options.Index().SetUnique(true),
	}

	studC.Indexes().CreateOne(ctx, index)
	ss = student_storage.NewStudentStorage(studC, ctx)
	sc = student_business.New(ss)
	server = gin.Default()
}

func InsertData(thread int) {
	var testStudC *mongo.Collection
	var testMongoClient *mongo.Client

	uriConn := "mongodb://" + USER + ":" + PASSWORD + "@" + HOST + ":" + PORT + "/" + DB_NAME + "?authSource=admin"
	option := options.Client().ApplyURI(uriConn)
	testMongoClient, err := mongo.Connect(ctx, option)
	if err != nil {
		log.Fatal("error while connecting with mongo", err)
	}

	err = testMongoClient.Ping(context.Background(), readpref.Primary())
	if err != nil {
		log.Fatal("error while trying to ping mongo", err)
	}

	fmt.Println("mongo connection established")
	mu.Lock()

	testStudC = testMongoClient.Database("student_dev_training").Collection("student")
	idInc, _ := testStudC.CountDocuments(ctx, bson.D{})

	randomStud := &student_model.Student{
		SID:    strconv.FormatInt(idInc, 10),
		Name:   utils.RandomName(),
		Class:  utils.RandomClass(),
		Gender: utils.RandomGender(),
	}

	testStudC.InsertOne(ctx, randomStud)

	mu.Unlock()

	wg.Done()

	defer testMongoClient.Disconnect(ctx)
}

func InsertConcurrency() {
	n := 10
	for i := 0; i < n; i++ {
		wg.Add(1)
		go InsertData(i)
	}
	wg.Wait()
}

func main() {
	defer mongoClient.Disconnect(ctx)

	basePath := server.Group("/v1")
	sc.RegisterUserRoutes(basePath)

	log.Fatal(server.Run(":5678"))
	//InsertConcurrency()
}
