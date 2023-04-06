package router

import (
	"go-fiber-crud/controller"

	"github.com/gofiber/fiber/v2"
)

func NewRouter(noteController *controller.NoteController) *fiber.App {
	router := fiber.New()

	router.Get("/healthchecker", func(c *fiber.Ctx) error {
		return c.Status(200).JSON(fiber.Map{
			"status":  "success",
			"message": "Welcome to Golang, Fiber, and Gorm",
		})
	})

	router.Route("/notes", func(router fiber.Router) {
		router.Post("/", noteController.Create)
		router.Get("", noteController.FindAll)
	})

	router.Route("/notes/:noteId", func(router fiber.Router) {
		router.Get("", noteController.FindById)
		router.Patch("", noteController.Update)
		router.Delete("", noteController.Delete)
	})

	return router
}
