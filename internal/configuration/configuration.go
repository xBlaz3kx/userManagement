package configuration

import (
	"github.com/kkyr/fig"
	"log"
)

type (
	Configuration struct {
		API   API   `fig:"api" validate:"required"`
		Mongo Mongo `fig:"mongo" validate:"required"`
	}

	API struct {
		Host string `fig:"host" default:"localhost"`
		Port int    `fig:"port" default:"8080"`
	}

	// Mongo configuration
	Mongo struct {
		Host       string `fig:"host" default:"localhost"`
		Username   string `fig:"username" validate:"required"`
		Password   string `fig:"password" validate:"required"`
		Port       int    `fig:"port" default:"27017"`
		Database   string `fig:"database" default:"test"`
		ReplicaSet string `fig:"replicaSet"`
	}
)

func GetConfiguration() *Configuration {
	var config Configuration

	err := fig.Load(&config,
		fig.File("config.yaml"),
		fig.Dirs("../../config", "/user-management/config"),
	)
	if err != nil {
		log.Fatal(err)
	}

	return &config
}
