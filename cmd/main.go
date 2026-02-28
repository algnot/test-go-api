package main

import (
	"web-server-101/di/config"
	"web-server-101/di/database"
	"web-server-101/di/server"
)

func main() {
	cfg := config.GetConfig()

	if cfg.Server.Service == "server" {
		err := server.InitApiServer()
		if err != nil {
			panic(err)
		}
	} else if cfg.Server.Service == "migrator" {
		db, err := database.InitDatabase()
		if err != nil {
			panic(err)
		}
		err = database.MigrateDatabase(db)
	}
}
