package main

import (
	"log"

	"github.com/warnawski/space-evolution/internal/configurate"
	"github.com/warnawski/space-evolution/pkg/logger"
)

func main() {

	logger.ConfigureLogger()

	conf := configurate.NewConf("../../examples/configuration/config.yaml")
	if conf == nil {
		return
	}

	err := conf.LoadConfig()
	if err != nil {
		return
	}
	log.Printf("Server configuration loaded: server_name: %s", conf.Data.ServerName)

}
