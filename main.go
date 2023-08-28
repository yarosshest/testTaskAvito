package main

import (
	"testTaskAvito/api"
	"testTaskAvito/db"
)

// @title test task avito API
// @version 1.0
// @description This is a test task avito server.

// @host localhost:8080
// @BasePath /

func main() {
	db.DbInit()
	api.RunApi()
}
