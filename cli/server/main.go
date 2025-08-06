package main

import (
	"github.com/Segian/little-sender/internal/api"
)

func main() {
	router := api.NewRouter()
	if err := router.Run(":8080"); err != nil {
		panic("Failed to start server: " + err.Error())
	}
}
