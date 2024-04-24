package models

import (
	"fmt"
	"time"

	"gorm.io/gorm"
)

type Users struct {
	gorm.Model
	EmpID     string `json:"employee_id"`
	Name      string `json:"name"`
	Lastname  string `json:"lastname"`
	Birthdate string `json:"birthdate"`
	Age       int    `json:"age"`
	Email     string `json:"email"`
	Tel       string `json:"tel"`
}

type UsersRes struct {
	EmpID     string `json:"employee_id"`
	Name      string `json:"name"`
	Lastname  string `json:"lastname"`
	Birthdate string `json:"birthdate"`
	Age       int    `json:"age"`
	Email     string `json:"email"`
	Tel       string `json:"tel"`
	Type      string `json:"type"`
}

func (u *Users) CalculateAge() {
	birthDate, err := time.Parse("2006-01-02", u.Birthdate)
	if err != nil {
		fmt.Println("Error parsing birthdate:", err)
		return
	}
	now := time.Now()
	years := now.Year() - birthDate.Year()
	if now.Month() < birthDate.Month() || (now.Month() == birthDate.Month() && now.Day() < birthDate.Day()) {
		years--
	}
	u.Age = years
}
