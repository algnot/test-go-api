package main

import (
	"log"
	"web-server-101/di/config"
	"web-server-101/di/database"
	"web-server-101/di/server"
)

func main() {
	cfg := config.GetConfig()
	log.Printf("server is running")

	if cfg.Server.Service == "server" {
		log.Printf("server is started")
		err := server.InitApiServer()
		if err != nil {
			panic(err)
		}
	} else if cfg.Server.Service == "migrator" {
		log.Printf("Migrator is started")
		db, err := database.InitDatabase()
		if err != nil {
			log.Printf("InitDatabase error")
			panic(err)
		}
		err = database.MigrateDatabase(db)
		if err != nil {
			log.Printf("MigrateDatabase error")
			panic(err)
		}
	}
}
