package main

import (
	"log"

	"github.com/behdadzabihi/car-buying-selling-system/src/api"
	"github.com/behdadzabihi/car-buying-selling-system/src/config"
	"github.com/behdadzabihi/car-buying-selling-system/src/data/cache"
	"github.com/behdadzabihi/car-buying-selling-system/src/data/db"
)

func main() {
	cfg := config.GetConfig()
	api.InitServer(cfg)
	defer cache.CloseRedis()
	err := cache.InitRedis(cfg)
	if err != nil {
		log.Fatal(err)
	}
	err = db.InitDb(cfg)
	if err != nil {
		log.Fatal(err)
	}
	defer db.CloseDb()

}
