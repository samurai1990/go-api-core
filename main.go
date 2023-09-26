package main

import (
	"core_api/api"
	"core_api/database"
	errorCode "core_api/errors"
	"core_api/utils"
	"errors"
	"log"
	"os"
	"strings"
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
				if errors.Is(err, errorCode.ErrDuplicateKey) {
					log.Println(strings.Split(err.Error(), ",")[0])
					return
				} else {
					log.Fatalln(err)
				}
			} else {
				return
			}
		}
	}

	server := api.NewServer(conf.BIND_HOST, conf.BIND_PORT)
	r := server.Setup()
	if err := r.Run(server.ListenAddr); err != nil {
		log.Fatalf("not running with error: %s", err.Error())
	}
}
