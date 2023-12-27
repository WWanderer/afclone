package routes

import (
	"github.com/WWanderer/afclone/controllers"
	"github.com/gofiber/fiber/v2"
)

func Form1Routes(route fiber.Router) {
	route.Get("/", controllers.CreateForm1)
	route.Get("/:formid/context", controllers.GetForm1Context)
	route.Patch("/:formid/context", controllers.PatchForm1Context)
}