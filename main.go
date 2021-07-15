package main

import (
	"log"
	"rest-api/database"
	"rest-api/server"
)

// @title rest-api DOC
// @version beta
// @description This is the rest-api documentation
// @host localhost:7010

// @securityDefinitions.apikey Authorization
// @in header
// @name Authorization
func main() {
	err := database.Migrate()
	if err != nil {
		log.Println("Failed to migrate database : ", err)
	}
	server.Init()
}
