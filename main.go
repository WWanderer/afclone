package main

import (
	"fmt"
	"log"

	"github.com/WWanderer/afclone/models/common"
	"github.com/WWanderer/afclone/models/form1"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/gofiber/storage/mongodb/v2"
	"github.com/gofiber/template/html/v2"
)

type Duration struct {
	Value int
}

type Ccm2DisplayField struct {
	DisplayValue string
	Ccm2         common.Ccm2
}

func main() {
	// translations := translations.GetTranslations()

	store:= mongodb.New(mongodb.Config{
		ConnectionURI: "mongodb://admin:password@127.0.0.1:27017",
		Database: "toto",
		Collection: "totoStorage",

	})
	fmt.Println(store.Conn())

	engine := html.New("resources/views/form1", ".html")
	engine.Reload(true)
	engine.Debug(true)

	app := fiber.New(fiber.Config{
		Views: engine,
	})
	app.Use(logger.New())

	app.Get("/context", func(c *fiber.Ctx) error {
		form1Context := form1.Context{} // TODO or load from db
		// TODO inputs to be loaded from config file
		return c.Render("context", fiber.Map{"context": form1Context,
			"durations":        []Duration{{Value: 1}, {Value: 2}, {Value: 3}},
			"nationalAgencies": []Ccm2DisplayField{{DisplayValue: "BE01", Ccm2: 100}, {DisplayValue: "FR01", Ccm2: 101}, {DisplayValue: "DE01", Ccm2: 103}},
			"languages":        []Ccm2DisplayField{{DisplayValue: "EN", Ccm2: 200}, {DisplayValue: "FR", Ccm2: 201}, {DisplayValue: "DE", Ccm2: 202}},
		})
	})

	app.Patch("/context", func(c *fiber.Ctx) error {
		postedContext := new(form1.ContextDTO)
		if err := c.BodyParser(postedContext); err != nil {
			return err
		}
		context := postedContext.Map() // save
		fmt.Println(postedContext, context)


		return c.Render("context", fiber.Map{"context": context,
			"durations":        []Duration{{Value: 1}, {Value: 2}, {Value: 3}},
			"nationalAgencies": []Ccm2DisplayField{{DisplayValue: "BE01", Ccm2: 100}, {DisplayValue: "FR01", Ccm2: 101}, {DisplayValue: "DE01", Ccm2: 103}},
			"languages":        []Ccm2DisplayField{{DisplayValue: "EN", Ccm2: 200}, {DisplayValue: "FR", Ccm2: 201}, {DisplayValue: "DE", Ccm2: 202}},
		})
	})

	log.Fatal(app.Listen("localhost:3000"))
}
