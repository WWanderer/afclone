package routes

import (
	"github.com/WWanderer/afclone/controllers"
	"github.com/gofiber/fiber/v2"
)

func Form1Routes(route fiber.Router) {
	route.Get("/", controllers.CreateForm1)

	route.Get("/:formid/context", controllers.GetForm1Context)
	route.Patch("/:formid/context", controllers.PatchForm1Context)

	route.Get("/:formid/participating_organisations", controllers.GetForm1ParticipatingOrganisations)
	route.Patch("/:formid/participating_organisations", controllers.PatchForm1ParticipatingOrganisations)

	route.Get("/:formid/activities", controllers.GetForm1Activities)
	route.Patch("/:formid/activities", controllers.PatchForm1Activities)

	route.Get("/:formid/annexes", controllers.GetForm1Annexes)
	route.Patch("/:formid/annexes", controllers.PatchForm1Annexes)
}