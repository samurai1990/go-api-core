package main

import (
	"core_api/api"
	"core_api/utils"
	"log"

)

func main() {

	conf := utils.NewConfig()
	conf.LoadConfig(".")

	// store := database.NewCrudHandler()
	server := api.NewServer(conf.BIND_HOST, conf.BIND_PORT)
	log.Fatal(server.Start())

}
