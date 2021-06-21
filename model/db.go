package model

import (
	"University-Information-Website/utils"
	"context"
	"fmt"
	"log"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var client *mongo.Client = nil
var db *mongo.Database = nil

func InitDb() {
	var err error
	url := utils.Db + "://" + utils.Dbuser + ":" + utils.DbPassWord + "@" + utils.Dbhost + ":" + utils.DbPort
	// Set client options
	println(url)
	clientOptions := options.Client().ApplyURI(url)
	// Connect to MongoDB
	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()
	client, err = mongo.Connect(ctx, clientOptions)
	if err != nil || client == nil {
		fmt.Println("database connect error")
		log.Fatal("database connect error", err)
	}
	// Check the connection
	err = client.Ping(context.TODO(), nil)
	if err != nil {
		fmt.Println("ping connect error")
		log.Fatal("ping connect error", err)
	}

	db = client.Database(utils.DbName)
	if db == nil {
		fmt.Println("switch database error")
		log.Fatal("switch database error", err)
	}

	fmt.Println("connected to MongoDB!")
}

func Close() {
	if client != nil {
		err := client.Disconnect(context.TODO())
		if err != nil {
			fmt.Println("close MongoDB error")
			log.Fatal("close MongoDB error", err)
		}
		fmt.Println("connection to MongoDB closed")
	} else {
		fmt.Println("not need close")
	}
}
