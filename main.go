package main

import (
	"log"

	"github.com/WWanderer/afclone/database/dbconfig"
	"github.com/WWanderer/afclone/routes"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/gofiber/template/html/v2"
)

type Duration struct {
	Value int
}

func main() {
	// translations := translations.GetTranslations()

	dbconfig.ConnectDB()

	engine := html.New("resources/views", ".html")
	engine.Reload(true)
	// engine.Debug(true)

	app := fiber.New(fiber.Config{
		Views: engine,
	})
	app.Use(logger.New())

	app.Get("/", func(c *fiber.Ctx) error {
		return c.Render("index", nil)
	})

	routes.Form1Routes(app.Group("form1"))

	log.Fatal(app.Listen("localhost:3000"))
}
