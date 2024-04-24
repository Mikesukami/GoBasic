package controllers

import (
	"go-workshop/database"
	m "go-workshop/models"

	"github.com/gofiber/fiber/v2"
)

func HelloTest(c *fiber.Ctx) error {
	return c.SendString("Hello, World!")
}

func GetUsers(c *fiber.Ctx) error {
	db := database.DBConn
	var users []m.Users

	db.Find(&users)
	return c.JSON(users)
}

func AddUser(c *fiber.Ctx) error {
	db := database.DBConn
	user := new(m.Users)

	if err := c.BodyParser(user); err != nil {
		return c.Status(400).SendString(err.Error())
	}

	user.CalculateAge()

	db.Create(&user)
	return c.JSON(user)
}

func UpdateUser(c *fiber.Ctx) error {
	db := database.DBConn
	id := c.Params("id")
	user := new(m.Users)

	if err := c.BodyParser(user); err != nil {
		return c.Status(400).SendString(err.Error())
	}

	var userDB m.Users
	db.First(&userDB, id)
	if userDB.EmpID == "" {
		return c.Status(404).SendString("User not found")
	}

	user.CalculateAge()
	db.Model(&userDB).Updates(&user)
	return c.JSON(userDB)
}

func RemoveUser(c *fiber.Ctx) error {
	db := database.DBConn
	id := c.Params("id")
	var user m.Users

	result := db.Delete(&user, id)

	if result.RowsAffected == 0 {
		return c.Status(404).SendString("User not found")
	}

	return c.SendString("User deleted successfully")
}

func GetUserJson(c *fiber.Ctx) error {
	db := database.DBConn
	var users []m.Users

	db.Find(&users)

	var dataResults []m.UsersRes
	sum_genZ := 0
	sum_genY := 0
	sum_genX := 0
	sum_genBoomer := 0
	sum_gi_gen := 0

	for _, v := range users {
		typeStr := ""
		if v.Age >= 0 && v.Age <= 23 {
			typeStr = "Gen Z"
			sum_genZ++
		} else if v.Age >= 24 && v.Age <= 41 {
			typeStr = "Gen Y"
			sum_genY++
		} else if v.Age >= 42 && v.Age <= 56 {
			typeStr = "Gen X"
			sum_genX++
		} else if v.Age >= 57 && v.Age <= 75 {
			typeStr = "Boomer"
			sum_genBoomer++
		} else {
			typeStr = "GI Gen"
			sum_gi_gen++
		}
		u := m.UsersRes{
			EmpID:     v.EmpID,
			Name:      v.Name,
			Lastname:  v.Lastname,
			Birthdate: v.Birthdate,
			Age:       v.Age,
			Email:     v.Email,
			Tel:       v.Tel,
			Type:      typeStr,
		}
		dataResults = append(dataResults, u)
	}

	type ResultData struct {
		Data         []m.UsersRes `json:"data"`
		Count        int          `json:"count"`
		SumGenZ      int          `json:"sum_genZ"`
		SumGenY      int          `json:"sum_genY"`
		SumGenX      int          `json:"sum_genX"`
		SumGenBoomer int          `json:"sum_genBoomer"`
		SumGiGen     int          `json:"sum_gi_gen"`
	}

	r := ResultData{
		Data:         dataResults,
		Count:        len(dataResults),
		SumGenZ:      sum_genZ,
		SumGenY:      sum_genY,
		SumGenX:      sum_genX,
		SumGenBoomer: sum_genBoomer,
		SumGiGen:     sum_gi_gen,
	}

	return c.Status(200).JSON(r)

}

func SearchUser(c *fiber.Ctx) error {
	//I want to search user by name or lastname or employee_id with search query
	db := database.DBConn
	search := c.Query("search")
	var users []m.Users

	db.Where("emp_id = ? OR name LIKE ? OR lastname LIKE ?", search, "%"+search+"%", "%"+search+"%").Find(&users)

	if len(users) == 0 {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{"message": "ไม่พบข้อมูล"})
	}

	return c.JSON(users)
}

// func GetSearchProfiles(c *fiber.Ctx) error {
//     query := c.Query("input")

//     var profiles []m.Profile
//     db := database.DBConn
//     db.Where("employee_id = ? OR name LIKE ? OR lastname LIKE ?", query, "%"+query+"%", "%"+query+"%").Find(&profiles)

//     if len(profiles) == 0 {
//         return c.Status(fiber.StatusNotFound).JSON(fiber.Map{"message": "ไม่พบข้อมูล"})
//     }
//     return c.Status(fiber.StatusOK).JSON(profiles)
// }
