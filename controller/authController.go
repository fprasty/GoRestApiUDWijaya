package controller

import (
	"fmt"
	"log"
	"regexp"
	"strings"

	"github.com/fprasty/GoRestApiUDWijaya/database"
	"github.com/fprasty/GoRestApiUDWijaya/models"
	"github.com/gofiber/fiber/v2"
)

func validateEmail(email string) bool {
	Re := regexp.MustCompile(`[a-z0-9._%+\-]+@[a-z0-9._%+\-]+\.[a-z0-9._%+\-]`)
	return Re.MatchString(email)
}

func Register(c *fiber.Ctx) error {
	var data map[string]interface{}
	var userData models.User
	if err := c.BodyParser(&data); err != nil {
		fmt.Println("Register>Unable to parse body")
	}

	//Check jika password lebih dari 6
	if len(data["password"].(string)) <= 6 {
		c.Status(400)
		return c.JSON(fiber.Map{
			"massage": "Password must be greater than 6 character",
		})
	}
	//Check FirstName harus berisikan >=20
	if len(data["first_name"].(string)) >= 20 {
		c.Status(400)
		return c.JSON(fiber.Map{
			"massage": "Nama depan tidak boleh lebih dari 20 karakter",
		})
	}
	//Check LastName harus berisikan >=20
	if len(data["last_name"].(string)) >= 20 {
		c.Status(400)
		return c.JSON(fiber.Map{
			"massage": "Nama belakang tidak boleh lebih dari 20 karakter",
		})
	}

	//Check email whitespace
	if !validateEmail(strings.TrimSpace(data["email"].(string))) {
		c.Status(400)
		return c.JSON(fiber.Map{
			"massage": "Invalid email address",
		})
	}

	//Check if email already exist
	database.DB.Where("email=?", strings.TrimSpace(data["email"].(string))).First(&userData)
	if userData.Id != 0 {
		c.Status(400)
		return c.JSON(fiber.Map{
			"message": "Email already exist",
		})
	}

	user := models.User{
		FirstName: data["first_name"].(string),
		LastName:  data["last_name"].(string),
		Phone:     data["phone"].(string),
		Email:     strings.TrimSpace(data["email"].(string)),
	}

	//HashPassword
	user.SetPassword(data["password"].(string))
	//CreateUser
	err := database.DB.Create(&user)
	if err != nil {
		log.Println(err)
	}
	c.Status(200)
	return c.JSON(fiber.Map{
		"massage": "Akun berhasi dibuat",
	})

}
