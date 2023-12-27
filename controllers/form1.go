package controllers

import (
	"fmt"
	"log"
	"maps"
	"reflect"

	"github.com/WWanderer/afclone/database/collections/applications"
	"github.com/WWanderer/afclone/models/common"
	"github.com/WWanderer/afclone/models/form1"
	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
	"go.mongodb.org/mongo-driver/bson"
)

// TODO load from config or db
var contextInputs = map[string]interface{}{
	"durations":        []common.ValueLabel[int]{{Value: 1}, {Value: 2}, {Value: 3}},
	"nationalAgencies": []common.Ccm2DisplayField{{DisplayValue: "BE01", Ccm2: 100}, {DisplayValue: "FR01", Ccm2: 101}, {DisplayValue: "DE01", Ccm2: 103}},
	"languages":        []common.Ccm2DisplayField{{DisplayValue: "EN", Ccm2: 200}, {DisplayValue: "FR", Ccm2: 201}, {DisplayValue: "DE", Ccm2: 202}},
	"menuItems":        []common.ValueLabel[string]{{Value: "context", Label: "Context"}, {Value: "participating_organisations", Label: "Participating Organisations"}},
}

func BuildNavMenu(form1 form1.Form1) []common.ValueLabel[string] {
	menuItems := []common.ValueLabel[string]{}

	return menuItems
}

func CreateForm1(c *fiber.Ctx) error {
	formid := uuid.New()
	newForm1 := form1.NewForm1(formid.String())
	_, err := applications.InsertOne(newForm1)
	if err != nil {
		return fiber.NewError(fiber.StatusInternalServerError, "Error creating application", err.Error())
	}
	return c.Redirect("/form1/" + formid.String() + "/context")
}

func GetForm1Context(c *fiber.Ctx) error {
	formid, err := uuid.Parse(c.Params("formid"))
	if err != nil {
		return fiber.ErrBadRequest
	}

	var form1 = new(form1.Form1)
	err = applications.FindById(formid, form1).Decode(&form1) // TODO move decode to persistance layer
	if err != nil {
		return fiber.NewError(fiber.StatusInternalServerError, "Error while decoding application data", err.Error())
	}

	BuildNavMenu(*form1)

	bind := fiber.Map{"context": form1.Context.Value, "formid": formid}
	maps.Copy(bind, contextInputs)
	return c.Render("form1/context", bind)
}

func PatchForm1Context(c *fiber.Ctx) error {
	formid, err := uuid.Parse(c.Params("formid"))
	if err != nil {
		return fiber.ErrBadRequest
	}

	postedContext := new(form1.ContextDTO)
	if err := c.BodyParser(postedContext); err != nil {
		log.Println("toto", err)
		return err
	}
	form1Context := common.Control[form1.Context]{Value: *postedContext.Map()}

	update := bson.D{{
		Key:   "$set",
		Value: bson.D{{Key: "context", Value: form1Context}},
	}}

	applications.UpdateOrInsert(formid, update)

	bind := fiber.Map{"context": form1Context.Value, "formid": formid}
	maps.Copy(bind, contextInputs)
	return c.Render("form1/context", bind)
}
