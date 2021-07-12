package models

import "gorm.io/gorm"

type Appointment struct {
	gorm.Model
	patient     Patient
	doctor      Doctor
	description string
}
