package main

import (
	"awesomeProject/api"
	"awesomeProject/db"
)

func main() {
	db.DbInit()
	api.RunApi()
}
