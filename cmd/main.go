package main

import (
	"progmamun/go-backend-boilerplate/internal/config"
	"progmamun/go-backend-boilerplate/internal/server"
)

func main() {
	// load environment variables
	cfg := config.LoadEnv()

	// connect to the database
	db := config.ConnectDatabase(cfg)

	// start the server
	server.Start(db, cfg)
}