package repository

import (
	"context"
	"fmt"

	log "github.com/sirupsen/logrus"

	"github.com/walbety/go-workshop/user-service/cmd/canonical"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var Client mongo.Client

func StartConnectionMongo(c context.Context) (*mongo.Client, error) {

	clientOptions := options.Client().ApplyURI("mongodb://root:rootpassword@localhost:27017")

	client, err := mongo.Connect(context.Background(), clientOptions)
	if err != nil {
		log.Fatal(err)
	}

	err = client.Ping(context.Background(), nil)
	if err != nil {
		log.Fatal(err)
	}

	return client, nil
}

func InsertOne() {
	user1 := canonical.User{
		Name:  "wal1",
		Email: "wal2",
	}

	collection := Client.Database("testdb").Collection("user")
	insertResult, err := collection.InsertOne(context.Background(), user1)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Documento inserido com ID:", insertResult.InsertedID)
	// Consultar documentos
}
