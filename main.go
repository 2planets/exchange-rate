package main

import (
	"exchange/config"
	"exchange/modules"
	"exchange/routes"
	"flag"
	"log"

	"github.com/gin-gonic/gin"
)

var configPath = flag.String("config", "config.json", "Set config path and filename")

func main() {
	flag.Parse()

	// Load the configs settings from the specified file
	configs := config.GetConfig(*configPath)
	log.Printf("Listening on HOST: %s\n", configs.Host)

	if err := modules.LoadCurrencyConfig(configs.RatePath); err != nil {
		log.Println(err)
	}

	gin.SetMode(gin.ReleaseMode)
	router := routes.Router()
	router.Run(configs.Host)
}
