package handler

import "github.com/gofiber/fiber/v2"

func GetUser(c *fiber.Ctx) error {
	return c.SendStatus(fiber.StatusOK)
}
