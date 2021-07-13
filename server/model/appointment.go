package model

import "gorm.io/gorm"

type Appointment struct {
	gorm.Model
	DoctorName  string
	Registrants []*User `gorm:"many2many:user_appointments;"`
	Description string
	Capacity    uint
	IsFull      bool
}
