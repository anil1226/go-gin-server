package main

import (
	"go-http/api"
	"log"
)

func main() {
	err := api.Start()
	if err != nil {
		log.Fatalln("Failed to start server", err)
	}

}
