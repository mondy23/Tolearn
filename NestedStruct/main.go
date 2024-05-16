package main

import (
	"Go/models"
	"log"

	"github.com/gofiber/fiber/v2"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var (
	db    *gorm.DB
	dberr error
)

func main() {
	dsn := "host=localhost user=postgres password='Reymond23s' dbname=dbtest port=5432 sslmode=disable TimeZone=Asia/Shanghai"
	db, dberr = gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if dberr != nil {
		panic(dberr)
	}
	db.AutoMigrate(&models.User{}, &models.Profile{})
	app := fiber.New()
	app.Get("/", Nested)
	log.Fatal(app.Listen(":3000"))
}

func Nested(c *fiber.Ctx) error {
	var user models.User
	err := c.BodyParser(&user)
	if err != nil {
		panic(err)
	}
	db.Create(&user)
	// user := &models.User{
	// 	Username: "SaucyKills",
	// 	Password: "123345",
	// 	Profile: models.Profile{
	// 		Name: "Reymond",
	// 		Age:  25,
	// 	},
	// }
	return c.JSON(user)
}
