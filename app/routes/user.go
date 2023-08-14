// routes/user_route.go
package routes

import (
	"attendance/app/controllers"
	"github.com/gofiber/fiber/v2"
)

func UserRoute(app *fiber.App, userController *controllers.UserController) {
	app.Post("/user", userController.CreateUser)
	// app.Get("/user/:userId", userController.GetAUser)
	// app.Put("/user/:userId", userController.EditAUser)
	// app.Delete("/user/:userId", userController.DeleteAUser)
	app.Get("/api/v1/users", userController.GetAllUsers)
}
