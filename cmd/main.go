package main

import (
	"dooplik/avitoTestTask/pkg/common/config"
	"dooplik/avitoTestTask/pkg/common/storage"
	"dooplik/avitoTestTask/pkg/server/handlers"
	"log"

	"github.com/gofiber/fiber/v2"
)

func main() {
	c, err := config.LoadConfig()

    if err != nil {
        log.Fatalln("Failed at config", err)
    }

	h := db.InitPostgresDB(&c)
	err = h.InsertTagsAndFeatures()
	if err != nil {
		log.Fatalln("Failed to insert tags and features", err)
	}

	l, err := db.InitInMemoryDB(h)
	if err != nil {
		log.Fatalln("Failed init local db", err)
	}
	go l.UpdateInMemotyDB(h)

	app := fiber.New()

	handlers.RegisterRoutes(app, h, l)
    
	app.Listen(c.Port)
}