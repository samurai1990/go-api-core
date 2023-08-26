package main

import (
	"core_api/api"
	"core_api/database"
	"core_api/utils"
	"log"
	"os"
)

func main() {

	conf := utils.NewConfig()
	conf.LoadConfig(".")

	store := database.NewDBHandler()
	if err := store.DBConnection(); err != nil {
		log.Fatal("connention db error")
	}

	for _, arg := range os.Args[1:] {
		if arg == "createsuperuser" {
			db := database.NewUser()
			if err := db.CreateSuperUser(); err != nil {
				log.Println(err)
			}
		}
	}

	server := api.NewServer(conf.BIND_HOST, conf.BIND_PORT)
	log.Fatal(server.Start())

}
