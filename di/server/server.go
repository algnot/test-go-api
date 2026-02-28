package server

import (
	"log"
	"web-server-101/di/config"
	"web-server-101/di/database"
	router "web-server-101/service"

	"github.com/gofiber/fiber/v3"
)

func InitApiServer() error {
	cfg := config.GetConfig()
	app := fiber.New()

	app.Get("/", func(c fiber.Ctx) error {
		return c.SendString("server is running")
	})

	db, err := database.InitDatabase()
	if err != nil {
		panic(err)
	}

	router.InitRouter(app, db)

	log.Fatal(app.Listen(":" + cfg.Server.AppPort))
	return nil
}
