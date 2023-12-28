package controllers

import (
	"log"
	"maps"

	"github.com/WWanderer/afclone/database/collections/applications"
	"github.com/WWanderer/afclone/models/common"
	"github.com/WWanderer/afclone/models/form1"
	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
	"go.mongodb.org/mongo-driver/bson"
)

var menuItems = []common.ValueLabel[string]{
	{Value: "context", Label: "Context"},
	{Value: "participating_organisations", Label: "Participating Organisations"},
	{Value: "activities", Label: "Activities"},
	{Value: "annexes", Label: "Annexes"},
}

// TODO load from config or db
var contextInputs = map[string]interface{}{
	"durations":        []common.ValueLabel[int]{{Value: 1}, {Value: 2}, {Value: 3}},
	"nationalAgencies": []common.Ccm2DisplayField{{DisplayValue: "BE01", Ccm2: 100}, {DisplayValue: "FR01", Ccm2: 101}, {DisplayValue: "DE01", Ccm2: 103}},
	"languages":        []common.Ccm2DisplayField{{DisplayValue: "EN", Ccm2: 200}, {DisplayValue: "FR", Ccm2: 201}, {DisplayValue: "DE", Ccm2: 202}},
	"menuItems":        menuItems,
}

var participatingOrganisationsInputs = map[string]interface{}{
	"menuItems": menuItems,
}
var activitiesInputs = map[string]interface{}{
	"menuItems": menuItems,
}
var annexesInputs = map[string]interface{}{
	"menuItems": menuItems,
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

func GetForm1ParticipatingOrganisations(c *fiber.Ctx) error {
	bind := fiber.Map{}
	maps.Copy(bind, participatingOrganisationsInputs)
	return c.Render("form1/participating_organisations", bind)
}
func PatchForm1ParticipatingOrganisations(c *fiber.Ctx) error {
	bind := fiber.Map{}
	maps.Copy(bind, participatingOrganisationsInputs)
	return c.Render("form1/participating_organisations", bind)
}
func GetForm1Activities(c *fiber.Ctx) error {
	bind := fiber.Map{}
	maps.Copy(bind, activitiesInputs)
	return c.Render("form1/activities", bind)
}
func PatchForm1Activities(c *fiber.Ctx) error {
	bind := fiber.Map{}
	maps.Copy(bind, activitiesInputs)
	return c.Render("form1/activities", bind)
}
func GetForm1Annexes(c *fiber.Ctx) error {
	bind := fiber.Map{}
	maps.Copy(bind, annexesInputs)
	return c.Render("form1/annexes", bind)
}
func PatchForm1Annexes(c *fiber.Ctx) error {
	bind := fiber.Map{}
	maps.Copy(bind, annexesInputs)
	return c.Render("form1/annexes", bind)
}
