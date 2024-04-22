package routes

import (
	// เช็คความถูกต้องของข้อมูล
	"github.com/gofiber/fiber/v2" // สร้าง API
	"github.com/gofiber/fiber/v2/middleware/basicauth"

	c "go-fiber-test/controllers" // สร้าง middleware สำหรับการเข้ารหัสข้อมูล
)

func InetRoutes(app *fiber.App) {
	app.Use(basicauth.New(basicauth.Config{
		Users: map[string]string{
			"gofiber": "21022566",
		},
	}))

	// /api/v1
	api := app.Group("/api")
	v1 := api.Group("/v1")
	v1.Get("/", c.HelloTest)
	v1.Post("/", c.BodyParserTest)
	v1.Get("/user/:name", c.ParamsTest)
	v1.Get("/search", c.QueryTest)
	v1.Post("/inet", c.QueryTest)
	v1.Post("/valid", c.ValidTest)
	v1.Post("/fact/:num", c.FactorialTest)
	v1.Post("/register", c.RegisterUser)

	//CRUD dogs
	dog := v1.Group("/dog")
	dog.Get("", c.GetDogs)
	dog.Get("/filter", c.GetDog)  //search query
	dog.Get("/del", c.GetDelDogs) //7.0.2
	dog.Get("/con", c.GetDogsCon) // 7.1
	dog.Get("/json", c.GetDogsJson)
	dog.Post("/", c.AddDog)
	dog.Put("/:id", c.UpdateDog)
	dog.Delete("/:id", c.RemoveDog)

	//CRUD company
	company := v1.Group("/company")
	company.Get("", c.GetCompany)
	company.Post("/", c.AddCompany)
	company.Put("/:id", c.UpdateCompany)
	company.Delete("/:id", c.RemoveCompany)

	v2 := api.Group("/v2")
	v2.Get("/", c.HelloTestV2)

	v3 := api.Group("/v3")
	v3.Post("/", c.NameToAscii)
}
