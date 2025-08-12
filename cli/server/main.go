package main

import (
	"github.com/Segian/little-sender/internal/api"
	"github.com/Segian/little-sender/internal/config"
)

func main() {
	if err := config.InitDatabase(); err != nil {
		panic("Failed to initialize database: " + err.Error())
	}
	router := api.NewRouter()
	if err := router.Run(":8080"); err != nil {
		panic("Failed to start server: " + err.Error())
	}

}
