package models

type User struct {
	FirstName string
	LastName  string
	Age       int
	Email     string `gorm:"primaryKey"`
	Username  string
	Password  string
}

type Doctor struct {
	User
	Registrants []string
}

type Administrator struct {
	User
}

type Patient struct {
	User
}
