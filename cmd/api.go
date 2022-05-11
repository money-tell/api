package main

import (
	"log"

	"github.com/katalabut/money-tell-api/internal/api"
	"github.com/katalabut/money-tell-api/internal/config"
)

func main() {

	cfg, err := config.InitConfig()
	if err != nil {
		log.Fatalln(err.Error())
	}

	err = api.ServeApi(cfg)
	if err != nil {
		log.Fatalln(err.Error())
	}
}
