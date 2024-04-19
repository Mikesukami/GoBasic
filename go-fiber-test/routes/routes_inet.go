package routes

import (
	// เช็คความถูกต้องของข้อมูล
	"github.com/gofiber/fiber/v2" // สร้าง API
	"github.com/gofiber/fiber/v2/middleware/basicauth"

	"go-fiber-test/controllers" // สร้าง middleware สำหรับการเข้ารหัสข้อมูล
)

func InetRoutes(app *fiber.App) {
	app.Use(basicauth.New(basicauth.Config{
		Users: map[string]string{
			"john":  "doe",
			"admin": "123456",
		},
	}))

	// /api/v1
	api := app.Group("/api")
	v1 := api.Group("/v1")
	v1.Get("/", controllers.HelloTest)
	v1.Post("/", controllers.BodyParserTest)
	v1.Get("/user/:name", controllers.ParamsTest)
	v1.Get("/search", controllers.QueryTest)
	v1.Post("/inet", controllers.QueryTest)
	v1.Post("/valid", controllers.ValidTest)

	v2 := api.Group("/v2")
	v2.Get("/", controllers.HelloTestV2)
}
