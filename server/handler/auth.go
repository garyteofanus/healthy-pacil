package handler

import (
	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt"
	"golang.org/x/crypto/bcrypt"
	"healthy-pacil/config"
	"healthy-pacil/database"
	"healthy-pacil/model"
	"strconv"
	"time"
)

func Register(c *fiber.Ctx) error {
	user := new(model.User)
	user.Role = model.Patient
	if parseError := c.BodyParser(user); parseError != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": "Parsing error",
			"detail":  parseError.Error(),
		})
	}

	hashedPassword, passErr := bcrypt.GenerateFromPassword([]byte(user.Password), 14)
	if passErr != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": "Hashing error",
			"detail":  passErr.Error(),
		})
	}
	user.Password = string(hashedPassword)

	if databaseErr := database.DB.Create(&user).Error; databaseErr != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": "Database error",
			"detail":  databaseErr.Error(),
		})
	}

	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"message": "User created",
		"data":    user,
	})
}

func Login(c *fiber.Ctx) error {
	type input struct {
		username string
		password string
	}

	userInput := new(input)
	if parseError := c.BodyParser(userInput); parseError != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": "Parsing error",
			"detail":  parseError.Error(),
		})
	}

	var user model.User
	database.DB.Where("username = ?", userInput.username).First(&user)

	cmpErr := bcrypt.CompareHashAndPassword([]byte(userInput.password), []byte(user.Password))
	if cmpErr != nil {
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
			"message": "Invalid password",
			"detail":  "Check your password again",
		})
	}

	claims := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.StandardClaims{
		Issuer:    strconv.FormatUint(uint64(user.ID), 10),
		ExpiresAt: time.Now().Add(time.Hour * 24).Unix(),
	})

	token, jwtErr := claims.SignedString([]byte(config.Config("SECRET")))
	if jwtErr != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": "JWT Error",
			"detail":  jwtErr.Error(),
		})
	}

	cookie := fiber.Cookie{
		Name:     "jwt",
		Value:    token,
		Expires:  time.Now().Add(time.Hour * 24),
		HTTPOnly: true,
	}

	c.Cookie(&cookie)

	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"message": "Login success",
		"data":    user,
	})
}
