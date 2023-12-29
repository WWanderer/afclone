package controllers

import (
	"fmt"
	"maps"
	"strconv"

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
	"activityTypes": []common.Ccm2DisplayField{{DisplayValue: "Mobility Activity Partner Countries", Ccm2: 1001}, {DisplayValue: "Preparatory Visit", Ccm2: 1002}, {DisplayValue: "Mobility Activity Countries Not Associated", Ccm2: 1004}},
	"menuItems":     menuItems,
}
var annexesInputs = map[string]interface{}{
	"menuItems": menuItems,
}

func BuildNavMenu(form1 form1.Form1) []common.ValueLabel[string] {
	menuItems := []common.ValueLabel[string]{}

	return menuItems
}

func fetchForm1(formid string) (*form1.Form1, error) {
	uuid, err := uuid.Parse(formid)
	if err != nil {
		return nil, fmt.Errorf("parsing %s as uuid: %w", formid, err)
	}

	var f1 = new(form1.Form1)
	err = applications.FindById(uuid, &f1)
	if err != nil {
		return nil, fmt.Errorf("decoding form %s from db: %w", formid, err)
	}

	return f1, nil
}

func updateForm1(formid string, update bson.D) error {
	uuid, err := uuid.Parse(formid)
	if err != nil {
		return fmt.Errorf("parsing %s as uuid: %w", formid, err)
	}

	_, err = applications.UpdateOrInsert(uuid, update)
	if err != nil {
		return fmt.Errorf("updating context for %s: %w", formid, err)
	}
	return nil
}

func CreateForm1(c *fiber.Ctx) error {
	formid := uuid.New()
	newForm1 := form1.NewForm1(formid.String())
	_, err := applications.InsertOne(newForm1)
	if err != nil {
		return fmt.Errorf("creating form %s: %w", formid, err)
	}
	return c.Redirect("/form1/" + formid.String() + "/context")
}

func GetForm1Context(c *fiber.Ctx) error {
	form1, err := fetchForm1(c.Params("formid"))
	if err != nil {
		return err
	}
	bind := fiber.Map{"context": form1.Context.Value, "formid": c.Params("formid")}
	maps.Copy(bind, contextInputs)
	return c.Render("form1/context", bind)
}

func PatchForm1Context(c *fiber.Ctx) error {
	formid := c.Params("formid")

	postedContext := new(form1.ContextDTO)

	if err := c.BodyParser(postedContext); err != nil {
		return fmt.Errorf("parsing context of %s from request body: %w", formid, err)
	}

	form1Context := common.Control[form1.Context]{Value: *postedContext.Map()}
	update := bson.D{{
		Key:   "$set",
		Value: bson.D{{Key: "context", Value: form1Context}},
	}}
	err := updateForm1(formid, update)
	if err != nil {
		return fmt.Errorf("updating form %s: %w", formid, err)
	}
	bind := fiber.Map{"context": form1Context.Value, "formid": formid}
	maps.Copy(bind, contextInputs)
	return c.Render("form1/context", bind)
}

func GetForm1ParticipatingOrganisations(c *fiber.Ctx) error {
	form1, err := fetchForm1(c.Params("formid"))
	if err != nil {
		return err
	}
	bind := fiber.Map{"participatingOrganisations": form1.ParticipatingOrganisations.Value, "formid": c.Params("formid")}
	maps.Copy(bind, participatingOrganisationsInputs)
	return c.Render("form1/participating_organisations", bind)
}

func PatchForm1ParticipatingOrganisations(c *fiber.Ctx) error {
	bind := fiber.Map{}
	maps.Copy(bind, participatingOrganisationsInputs)
	return c.Render("form1/participating_organisations", bind)
}

func GetForm1Activities(c *fiber.Ctx) error {
	form1, err := fetchForm1(c.Params("formid"))
	if err != nil {
		return err
	}
	bind := fiber.Map{"activities": form1.Activities.Value, "formid": c.Params("formid")}
	maps.Copy(bind, activitiesInputs)
	return c.Render("form1/activities", bind)
}

func PatchForm1Activities(c *fiber.Ctx) error {
	formid := c.Params("formid")

	fmt.Printf("c.Body(): %v\n", c.Body())

	bind := fiber.Map{"formid": formid}
	maps.Copy(bind, activitiesInputs)
	return c.Render("form1/activities", bind)
}

func Form1AddActivity(c *fiber.Ctx) error {
	formid := c.Params("formid")

	f1, err := fetchForm1(formid)
	if err != nil {
		return err
	}
	activityList := f1.Activities.Value.ActivityList.Value
	newActivity := form1.Activity{Id: common.Control[int]{Value: len(activityList)}}
	activityList = append(activityList,newActivity)

	f1.Activities.Value.ActivityList.Value = activityList
	update := bson.D{{
		Key:   "$set",
		Value: bson.D{{Key: "activities.value.activity_list", Value: f1.Activities.Value.ActivityList}},
	}}
	err = updateForm1(formid, update)
	if err != nil {
		return fmt.Errorf("updating form %s: %w", formid, err)
	}

	bind := fiber.Map{"formid": formid, "activities": f1.Activities.Value}
	maps.Copy(bind, activitiesInputs)
	return c.Render("form1/activities", bind)
}

func Form1DeleteActivity(c *fiber.Ctx) error {
	formid := c.Params("formid")
	activityid, err := strconv.Atoi(c.Params("id"))
	if err != nil {
		return fmt.Errorf("invalid activity id %s: %s", c.Params("id"), err)
	}
	f1, err := fetchForm1(formid)
	if err != nil {
		return err
	}

	activityList := f1.Activities.Value.ActivityList.Value
	fmt.Printf("activityList: %v\n", len(activityList))
	if len(activityList) == 1 {
		f1.Activities.Value.ActivityList.Value = []form1.Activity{}
	} else {
		al := make([]form1.Activity, 0)
		al = append(al, activityList[:activityid]...)
		for _, activity := range activityList[activityid+1:] {
			activity.Id.Value = activity.Id.Value - 1
			al = append(al, activity)
		}
		f1.Activities.Value.ActivityList.Value = al
	}
	update := bson.D{{
		Key:   "$set",
		Value: bson.D{{Key: "activities.value.activity_list", Value: f1.Activities.Value.ActivityList}},
	}}

	err = updateForm1(formid, update)
	if err != nil {
		return fmt.Errorf("updating form %s: %w", formid, err)
	}

	bind := fiber.Map{"formid": formid, "activities": f1.Activities.Value}
	maps.Copy(bind, activitiesInputs)
	return c.Render("form1/activities", bind)

}

func GetForm1Annexes(c *fiber.Ctx) error {
	form1, err := fetchForm1(c.Params("formid"))
	if err != nil {
		return err
	}
	bind := fiber.Map{"annexes": form1.Annexes.Value, "formid": c.Params("formid")}
	maps.Copy(bind, annexesInputs)
	return c.Render("form1/annexes", bind)
}

func PatchForm1Annexes(c *fiber.Ctx) error {
	bind := fiber.Map{}
	maps.Copy(bind, annexesInputs)
	return c.Render("form1/annexes", bind)
}
