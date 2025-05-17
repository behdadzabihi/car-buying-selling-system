package main

import (
	"github.com/behdadzabihi/car-buying-selling-system/src/api"
	"github.com/behdadzabihi/car-buying-selling-system/src/config"
	"github.com/behdadzabihi/car-buying-selling-system/src/data/cache"
)





func main() {
	cfg:=config.GetConfig()
	api.InitServer(cfg)
	defer cache.CloseRedis()
	cache.InitRedis(cfg) 
}