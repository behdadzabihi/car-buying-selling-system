package main

import (

	"github.com/behdadzabihi/car-buying-selling-system/src/api"
	"github.com/behdadzabihi/car-buying-selling-system/src/config"
	"github.com/behdadzabihi/car-buying-selling-system/src/data/cache"
	"github.com/behdadzabihi/car-buying-selling-system/src/data/db"
	"github.com/behdadzabihi/car-buying-selling-system/src/pkg/logging"
)

// @securityDefinitions.apikey AuthBearer
// @in header
// @name Authorization
func main() {
	cfg := config.GetConfig()
	logger := logging.NewLogger(cfg)
	defer cache.CloseRedis()
	err := cache.InitRedis(cfg)
	if err != nil {
		logger.Fatal(logging.Redis, logging.Startup, err.Error(), nil)
	}
	err = db.InitDb(cfg)
	if err != nil {
		logger.Fatal(logging.Postgres, logging.Startup, err.Error(), nil)
	}
	defer db.CloseDb()
	api.InitServer(cfg)
}
