package main

import (
	"fmt"
	"go-workshop/database"
	"go-workshop/routes"

	m "go-workshop/models"

	"github.com/gofiber/fiber/v2"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func initDatabase() {
	dsn := fmt.Sprintf(
		"%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=true&loc=Local",
		"root",
		"",
		"127.0.0.1",
		"3306",
		"go_fiber_ws",
	)
	var err error
	database.DBConn, err = gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		panic(err)
	}
	fmt.Println("Database connected!")
	database.DBConn.AutoMigrate(&m.Users{})
}

func main() {
	app := fiber.New()
	initDatabase()
	routes.GoWorkshopRoutes(app)
	app.Listen(":3000")
}
