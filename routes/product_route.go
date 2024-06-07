package routes

import (
	"github.com/lmhuong711/go-go-be/controllers"

	"github.com/gofiber/fiber/v2"
)

func SetupRoutes(api fiber.Router) {
	apiStudent := api.Group("/students")
	{
		apiStudent.Post("/", controllers.CreateStudent)
		apiStudent.Get("/", controllers.GetStudents)
		apiStudent.Get("/:id", controllers.GetStudent)
		apiStudent.Put("/:id", controllers.UpdateStudent)
		apiStudent.Delete("/:id", controllers.DeleteStudent)
	}
}
