package main

import (
	"fmt"
	"log"

	"github.com/Harsh-710/hospital-management/cmd/api"
	"github.com/Harsh-710/hospital-management/configs"
	"github.com/Harsh-710/hospital-management/db"
)

func main() {

	db.ConnectDB()

	database := db.GetDB()

	server := api.NewAPIServer(fmt.Sprintf(":%s", configs.Envs.Port), database)
	if err := server.Run(); err != nil {
		log.Fatal(err)
	}
}
