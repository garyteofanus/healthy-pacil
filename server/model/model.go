package model

import "gorm.io/gorm"

type User struct {
	firstName string
	lastName  string
	age       int
	email     string `gorm:"primaryKey"`
	username  string
	password  string
}

type Doctor struct {
	User
	registrants []string
}

type Appointment struct {
	gorm.Model
	patient     Patient
	doctor      Doctor
	description string
}

type Administrator struct {
	User
}

type Patient struct {
	User
}
