package db

import (
	"fmt"
	"github.com/kamva/mgm/v3"
	"github.com/xBlaz3kx/userManagementExample/internal/configuration"
	"go.mongodb.org/mongo-driver/mongo/options"
	"log"
)

func Connect(mongo configuration.Mongo) {
	databaseConnection := fmt.Sprintf(
		"mongodb://%s:%s@%s:%d/%s",
		mongo.Username,
		mongo.Password,
		mongo.Host,
		mongo.Port,
		mongo.Database,
	)

	opts := options.Client().ApplyURI(databaseConnection)
	if mongo.ReplicaSet != "" {
		opts.SetReplicaSet(mongo.ReplicaSet)
	}

	log.Println("Connecting to the database..")
	//try to connect to the database
	err := mgm.SetDefaultConfig(nil, mongo.Database, opts)
	if err != nil {
		log.Fatalf("Cannot connect to database:%v", err)
	}
}
