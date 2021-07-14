package handler

import (
	"fmt"
	"github.com/gofiber/fiber/v2"
	"healthy-pacil/database"
	"healthy-pacil/model"
)

func isValidUserPermission(user model.User, role model.Role) bool {
	if user.Role == role {
		return true
	}
	return false
}

func GetAllAppointments(c *fiber.Ctx) error {
	var appointments []model.Appointment
	database.DB.Find(&appointments)

	return c.Status(fiber.StatusOK).JSON(&appointments)
}

func ApplyAppointment(c *fiber.Ctx) error {
	appointmentId := c.Params("id")

	var appointment model.Appointment
	appointmentResult := database.DB.Where("id = ?", appointmentId).First(&appointment)

	if appointmentResult.Error != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": "Database error",
			"detail":  "Error getting appointment",
		})
	}

	claims, claimsErr := GetUserClaims(c)
	if claimsErr != nil {
		return claimsErr
	}

	var user model.User
	userResult := database.DB.First(&user, claims.StandardClaims.Issuer)

	if userResult.Error != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": "Database error",
			"detail":  "Error getting user",
		})
	}

	if appointment.IsFull || uint(len(appointment.Registrants)) == appointment.Capacity {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": "Appointment error",
			"detail":  "Appointment capacity is already full",
		})
	} else {
		appointment.Registrants = append(appointment.Registrants, &user)
	}

	return c.Status(fiber.StatusOK).JSON(&appointment)
}

func CancelAppointment(c *fiber.Ctx) error {
	appointmentId := c.Params("id")

	var appointment model.Appointment
	appointmentResult := database.DB.Where("id = ?", appointmentId).First(&appointment)
	if appointmentResult.Error != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": "Database error",
			"detail":  "Error getting appointment",
		})
	}

	//claims, claimsErr := GetUserClaims(c)
	//if claimsErr != nil {
	//	return claimsErr
	//}
	//
	//appointment.Registrants =

	return c.Status(fiber.StatusOK).JSON("")
}

func CreateAppointment(c *fiber.Ctx) error {
	appointment := new(model.Appointment)
	if parseError := c.BodyParser(appointment); parseError != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": "Parsing error",
			"detail":  parseError.Error(),
		})
	}

	database.DB.Create(appointment)

	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"message": "Appointment created",
		"data":    appointment,
	})
}

func UpdateAppointment(c *fiber.Ctx) error {
	appointmentId := c.Params("id")

	result := database.DB.Delete(&model.Appointment{}, appointmentId)

	if result.RowsAffected == 0 {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": "Update failed",
			"detail": fmt.Sprintf("Appointment with id %s is not found",
				appointmentId),
		})
	}

	appointment := new(model.Appointment)
	if parseError := c.BodyParser(appointment); parseError != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": "Parsing error",
			"detail":  parseError.Error(),
		})
	}

	database.DB.Save(appointment)

	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"message": "Appointment updated",
		"data":    appointment,
	})
}

func DeleteAppointment(c *fiber.Ctx) error {
	appointmentId := c.Params("id")

	result := database.DB.Delete(&model.Appointment{}, appointmentId)

	if result.RowsAffected == 0 {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": "Deletion failed",
			"detail": fmt.Sprintf("Appointment with id %s is not found",
				appointmentId),
		})
	}

	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"message": "Appointment deleted",
		"data":    result,
	})
}

func GetPatientsFromAppointment(c *fiber.Ctx) error {
	appointmentId := c.Params("id")

	var appointment model.Appointment
	appointmentResult := database.DB.Where("id = ?", appointmentId).First(&appointment)
	if appointmentResult.Error != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": "Database error",
			"detail":  "Error getting appointment",
		})
	}

	return c.Status(fiber.StatusOK).JSON(appointment.Registrants)
}
