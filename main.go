package main

import (
	"context"
	"log"
	"time"

	"github.com/WWanderer/afclone/config/db"
	"github.com/WWanderer/afclone/models/common"
	"github.com/WWanderer/afclone/models/form1"
	"github.com/google/uuid"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/gofiber/fiber/v2/middleware/session"
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

	db.ConnectDB()
	sessionStore := session.New()

	engine := html.New("resources/views", ".html")
	engine.Reload(true)
	engine.Debug(true)

	app := fiber.New(fiber.Config{
		Views: engine,
	})
	app.Use(logger.New())

	app.Get("/", func(c *fiber.Ctx) error {
		return c.Render("index", nil)
	})

	app.Get("/form1", func(c *fiber.Ctx) error {
		sess, err := sessionStore.Get(c)
		if err != nil {
			log.Panic(err)
		}
		if err := sess.Save(); err != nil {
			log.Panic(err)
		}
		formid := uuid.New()
		return c.Redirect("/form1/" + formid.String() + "/context")
	})

	app.Get("form1/:formid/context", func(c *fiber.Ctx) error {
		form1Context := form1.Context{} // TODO or load from db
		// TODO inputs to be loaded from config file
		return c.Render("form1/context", fiber.Map{"context": form1Context,
			"formid": c.Params("formid"),
			"durations":        []Duration{{Value: 1}, {Value: 2}, {Value: 3}},
			"nationalAgencies": []Ccm2DisplayField{{DisplayValue: "BE01", Ccm2: 100}, {DisplayValue: "FR01", Ccm2: 101}, {DisplayValue: "DE01", Ccm2: 103}},
			"languages":        []Ccm2DisplayField{{DisplayValue: "EN", Ccm2: 200}, {DisplayValue: "FR", Ccm2: 201}, {DisplayValue: "DE", Ccm2: 202}},
		})
	})

	app.Patch("form1/:formid/context", func(c *fiber.Ctx) error {
		postedContext := new(form1.ContextDTO)
		if err := c.BodyParser(postedContext); err != nil {
			log.Println(err)
			return err
		}
		formContext := postedContext.Map()

		applicationsCollection := db.Instance.DB.Collection("applications")
		ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
		defer cancel()
		applicationsCollection.InsertOne(ctx, formContext)

		return c.Render("form1/context", fiber.Map{"context": formContext,
			"formid":           c.Params("formid"),
			"durations":        []Duration{{Value: 1}, {Value: 2}, {Value: 3}},
			"nationalAgencies": []Ccm2DisplayField{{DisplayValue: "BE01", Ccm2: 100}, {DisplayValue: "FR01", Ccm2: 101}, {DisplayValue: "DE01", Ccm2: 103}},
			"languages":        []Ccm2DisplayField{{DisplayValue: "EN", Ccm2: 200}, {DisplayValue: "FR", Ccm2: 201}, {DisplayValue: "DE", Ccm2: 202}},
		})
	})

	log.Fatal(app.Listen("localhost:3000"))
}
