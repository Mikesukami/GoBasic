package routes

import (
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/basicauth"

	c "go-workshop/controllers"
)

func GoWorkshopRoutes(app *fiber.App) {
	api := app.Group("/api")
	v1 := api.Group("/v1")
	v1.Get("/", c.HelloTest)

	//CRUD users
	user := v1.Group("/user")
	user.Get("/", c.GetUsers)
	user.Post("/", Auth(), c.AddUser)
	user.Put("/:id", Auth(), c.UpdateUser)
	user.Delete("/:id", Auth(), c.RemoveUser)
	user.Get("/json", Auth(), c.GetUserJson)
	user.Get("/search", Auth(), c.SearchUser)
}

func Auth() fiber.Handler {
	return basicauth.New(basicauth.Config{
		Users: map[string]string{
			"testgo": "23012023",
		},
	})
}
