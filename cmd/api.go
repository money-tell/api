package main

import (
	"log"

	"github.com/katalabut/money-tell-api/app/api"
	"github.com/katalabut/money-tell-api/app/config"
)

func main() {
	log.Println("Config initialising")

	cfg, err := config.InitConfig("MoneyTell")
	if err != nil {
		log.Fatalln(err.Error())
	}

	log.Println("Config initialised")

	err = api.Run(cfg)
	if err != nil {
		log.Fatalln(err.Error())
	}
}
