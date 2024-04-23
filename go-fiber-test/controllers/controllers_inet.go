package controllers

import (
	"log"
	"strconv"
	"strings"

	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"

	m "go-fiber-test/models"

	"go-fiber-test/database"

	"regexp"
)

func HelloTest(c *fiber.Ctx) error {
	return c.SendString("Hello, World!")

}

func HelloTestV2(c *fiber.Ctx) error {
	return c.SendString("Hello, World! 2")

}

func BodyParserTest(c *fiber.Ctx) error {

	p := new(m.Person)

	if err := c.BodyParser(p); err != nil {
		return c.Status(500).SendString("Internal Server Error")
	}

	log.Println(p.Name) // john
	log.Println(p.Pass) // doe
	str := p.Name + p.Pass
	return c.JSON(str)
}

func ParamsTest(c *fiber.Ctx) error {

	str := "hello ==> " + c.Params("name")
	return c.JSON(str)
}

func QueryTest(c *fiber.Ctx) error {
	a := c.Query("search")
	str := "my search is  " + a

	return c.JSON(str)
}

func ValidTest(c *fiber.Ctx) error {
	//Connect to database

	user := new(m.User)
	if err := c.BodyParser(&user); err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": err.Error(),
		})
	}
	validate := validator.New()
	errors := validate.Struct(user)
	if errors != nil {
		return c.Status(fiber.StatusBadRequest).JSON(errors.Error())
	}

	return c.JSON(user)
}

func FactorialTest(c *fiber.Ctx) error {

	n := c.Params("num")        // รับค่าจาก params (n เป็น string)
	num, err := strconv.Atoi(n) // แปลงค่า n เป็น int ถ้ามี error จะ return ค่าเป็น nil
	fact := c.Params("num")     // รับค่าจาก params (fact เป็น string)
	if err != nil {             // ถ้ามี error จะ return ค่าเป็น nil
		return c.Status(fiber.StatusBadRequest).JSON("Invalid number")
	}

	// ลดค่า num ลง 1 แล้วคูณกับ i จนกว่า i จะเป็น 0
	for i := num - 1; i > 0; i-- {
		num *= i
	}

	// แปลงค่า num เป็น string
	n = strconv.Itoa(num)
	result := fact + "!=" + string(n)

	return c.JSON(result)
}

func NameToAscii(c *fiber.Ctx) error {
	//รับค่า Query Params ชื่อ name
	name := c.Query("tax_id")
	//แปลงชื่อเป็น ASCII
	ascii := ""
	for _, v := range name {
		ascii += strconv.Itoa(int(v)) + " "
	}

	return c.JSON(ascii)

}

func RegisterUser(c *fiber.Ctx) error {
	//Connect to database

	member := new(m.Member)
	if err := c.BodyParser(&member); err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": err.Error(),
		})
	}

	validate := validator.New()
	validate.RegisterValidation("username", validateUsername)
	validate.RegisterValidation("url", validateUrl)

	errors := validate.Struct(member)
	if errors != nil {
		return c.Status(fiber.StatusBadRequest).JSON(errors.Error())
	}

	return c.JSON(member)

}

func validateUsername(fl validator.FieldLevel) bool {
	username := fl.Field().String()
	if username == "" {
		return true // Allow empty string
	}
	regex := regexp.MustCompile(`^[a-zA-Z0-9_-]+$`)
	return regex.MatchString(username)
}

func validateUrl(fl validator.FieldLevel) bool {
	url := fl.Field().String()
	regex := regexp.MustCompile(`^[a-z0-9-]+$`)
	return regex.MatchString(url)
}

func DogIDGraterThan100(db *gorm.DB) *gorm.DB {
	return db.Where("dog_id > ?", 100)
}

func DogIDBetween50To100(db *gorm.DB) *gorm.DB {
	return db.Where("dog_id BETWEEN ? AND ?", 50, 100)
}

func GetDogs(c *fiber.Ctx) error {
	db := database.DBConn
	var dogs []m.Dogs

	db.Scopes(DogIDGraterThan100).Find(&dogs) //delelete = null
	return c.Status(200).JSON(dogs)
}

func GetDogsCon(c *fiber.Ctx) error {
	db := database.DBConn
	var dogs []m.Dogs

	db.Scopes(DogIDBetween50To100).Find(&dogs) //delelete = null
	return c.Status(200).JSON(dogs)
}

func GetDog(c *fiber.Ctx) error {
	db := database.DBConn
	search := strings.TrimSpace(c.Query("search"))
	var dog []m.Dogs

	result := db.Find(&dog, "dog_id = ?", search)

	// returns found records count, equals `len(users)
	if result.RowsAffected == 0 {
		return c.SendStatus(404)
	}
	return c.Status(200).JSON(&dog)
}

func AddDog(c *fiber.Ctx) error {
	//twst3
	db := database.DBConn
	var dog m.Dogs

	if err := c.BodyParser(&dog); err != nil {
		return c.Status(503).SendString(err.Error())
	}

	db.Create(&dog)
	return c.Status(201).JSON(dog)
}

func UpdateDog(c *fiber.Ctx) error {
	db := database.DBConn
	var dog m.Dogs
	id := c.Params("id")

	if err := c.BodyParser(&dog); err != nil {
		return c.Status(503).SendString(err.Error())
	}

	db.Where("id = ?", id).Updates(&dog)
	return c.Status(200).JSON(dog)
}

func RemoveDog(c *fiber.Ctx) error {
	db := database.DBConn
	id := c.Params("id")
	var dog m.Dogs

	result := db.Delete(&dog, id)

	if result.RowsAffected == 0 {
		return c.SendStatus(404)
	}

	return c.SendStatus(200)
}

func GetDogsJson(c *fiber.Ctx) error {
	db := database.DBConn
	var dogs []m.Dogs

	db.Find(&dogs) //10ตัว

	sum_red := 0
	sum_green := 0
	sum_pink := 0
	sum_no_color := 0

	var dataResults []m.DogsRes
	for _, v := range dogs { //1 inet 112 //2 inet1 113
		typeStr := ""
		if v.DogID >= 10 && v.DogID <= 50 {
			typeStr = "red"
			sum_red++
		} else if v.DogID >= 100 && v.DogID <= 150 {
			typeStr = "green"
			sum_green++
		} else if v.DogID >= 200 && v.DogID <= 250 {
			typeStr = "pink"
			sum_pink++
		} else {
			typeStr = "no color"
			sum_no_color++
		}

		d := m.DogsRes{
			Name:  v.Name,  //inet1
			DogID: v.DogID, //113
			Type:  typeStr, //green
		}
		dataResults = append(dataResults, d)
		// sumAmount += v.Amount
	}

	type ResultData struct {
		Data       []m.DogsRes `json:"data"`
		Name       string      `json:"name"`
		Count      int         `json:"count"`
		SumRed     int         `json:"sum_red"`
		SumGreen   int         `json:"sum_green"`
		SumPink    int         `json:"sum_pink"`
		SumNoColor int         `json:"sum_no_color"`
	}

	r := ResultData{
		Data:       dataResults,
		Name:       "golang-test",
		Count:      len(dogs), //หาผลรวม,
		SumRed:     sum_red,
		SumGreen:   sum_green,
		SumPink:    sum_pink,
		SumNoColor: sum_no_color,
	}
	return c.Status(200).JSON(r)
}

func GetDelDogs(ctx *fiber.Ctx) error {
	db := database.DBConn
	var dogs []m.Dogs
	// Retrieve all dogs including deleted records
	if err := db.Where("deleted_at IS NOT NULL").Find(&dogs).Error; err != nil {
		return ctx.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": "Failed to retrieve dogs",
			"error":   err.Error(),
		})
	}
	return ctx.JSON(dogs)
}

func GetCompany(ctx *fiber.Ctx) error {
	db := database.DBConn
	var company []m.Company
	// Retrieve all dogs including deleted records
	if err := db.Find(&company).Error; err != nil {
		return ctx.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": "Failed to retrieve company",
			"error":   err.Error(),
		})
	}
	return ctx.JSON(company)
}

func AddCompany(ctx *fiber.Ctx) error {
	db := database.DBConn
	var company m.Company
	if err := ctx.BodyParser(&company); err != nil {
		return ctx.Status(503).SendString(err.Error())
	}
	db.Create(&company)
	return ctx.Status(201).JSON(company)
}

func UpdateCompany(ctx *fiber.Ctx) error {
	db := database.DBConn
	var company m.Company
	id := ctx.Params("id")

	if err := ctx.BodyParser(&company); err != nil {
		return ctx.Status(503).SendString(err.Error())
	}

	db.Where("id = ?", id).Updates(&company)
	return ctx.Status(200).JSON(company)
}

func RemoveCompany(ctx *fiber.Ctx) error {
	db := database.DBConn
	id := ctx.Params("id")
	var company m.Company

	result := db.Delete(&company, id)

	if result.RowsAffected == 0 {
		return ctx.SendStatus(404)
	}

	return ctx.SendStatus(200)
}
