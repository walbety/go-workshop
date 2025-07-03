package main

import (
	"context"
	log "github.com/sirupsen/logrus"

	"github.com/walbety/go-workshop/user-service/cmd/config"
	"github.com/walbety/go-workshop/user-service/cmd/entrypoint/rest"
	"github.com/walbety/go-workshop/user-service/cmd/repository"
)

func main() {

	config.Initialize()
	mongoClient, err := repository.StartConnectionMongo(context.Background())
	if err != nil {
		log.Error("cant connect to mongodb")
	}

	rest.StartRest(mongoClient)

}
