package main

import (
	"api/api"
	"api/db"
)

// @title test task avito API
// @version 1.0
// @description This is a test task avito server.

// @host localhost:60122
// @BasePath /

func main() {
	db.DbInit()
	api.RunApi()
}
