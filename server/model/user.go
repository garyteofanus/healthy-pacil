package model

import "gorm.io/gorm"

type User struct {
	gorm.Model
	FirstName    string
	LastName     string
	Age          uint8
	Email        string `gorm:"unique_index;not null"`
	Username     string `gorm:"unique_index;not null"`
	Password     string
	Role         Role
	Appointments []*Appointment `gorm:"many2many:user_appointments;"`
}
