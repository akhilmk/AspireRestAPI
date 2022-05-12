package main

import (
	"log"

	api "github.com/akhilmk/GoRESTAPI/api"
	"github.com/akhilmk/GoRESTAPI/config"
	"github.com/akhilmk/GoRESTAPI/dbaccess"
)

func main() {
	log.Printf("Server Started !")

	config.LoadAppConfig()
	dbaccess.ConnectAndInitDB()
	api.StartService()
}
