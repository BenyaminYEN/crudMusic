package main

import (
	"crudMusic/config"
	"crudMusic/master"
)

func main() {
	db := config.ConnectDB()
	router := config.CreateRouter()
	master.Init(router, db)
	config.RunServer(router)
}
