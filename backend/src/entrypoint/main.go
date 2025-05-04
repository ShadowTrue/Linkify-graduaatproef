package main

import (
	dbservice "backend/src/db_service"
	"context"
	"fmt"
	"log"
	"time"

	"go.mongodb.org/mongo-driver/v2/mongo"
	"go.mongodb.org/mongo-driver/v2/mongo/options"
	//TODO "net/http"
)

func main() {
	fmt.Println("Hello World")

	db := createDbConnection()

	usrService := dbservice.CreateUserService(*db)
	
}


func createDbConnection() *mongo.Database{
	client, err := mongo.Connect(options.Client().ApplyURI("mongodb://localhost:27017/"))

	if err != nil{
		log.Fatal(err)
	}

	//this line handle the timeout, 
	// after an operation doesnt finish withing selected timeframe it cancel it.
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
    defer cancel()

	err = client.Ping(ctx,nil)

	if err != nil{
		log.Fatal("Could not ping the database", err)
	}

	db :=  client.Database("LinkifyDb")
	log.Println("Connection to database successfully established")
	return db
}
