package models

import "gorm.io/gorm"

type Person struct {
	Name string `json:"name"`
	Pass string `json:"pass"`
}

type User struct {
	Name     string `json:"name" validate:"required,min=3,max=32"`
	IsActive *bool  `json:"isactive" validate:"required"`
	Email    string `json:"email,omitempty" validate:"required,email,min=3,max=32"`
}

type Member struct {
	Email    string `json:"email,omitempty" validate:"required"`
	Username string `json:"username" validate:"username"`
	Password string `json:"password,omitempty" validate:"required,min=6,max=20"`
	LineID   string `json:"lineid"`
	Phone    string `json:"phone,omitempty" validate:"required,min=10,max=10"`
	BType    string `json:"btype"`
	SiteUrl  string `json:"siteurl" validate:"required,min=3,max=30,url"`
}

type Dogs struct {
	gorm.Model
	Name  string `json:"name"`
	DogID int    `json:"dog_id"`
}

type DogsRes struct {
	Name  string `json:"name"`
	DogID int    `json:"dog_id"`
	Type  string `json:"type"`
}

type Company struct {
	gorm.Model
	Name    string `json:"name"`
	Address string `json:"address"`
	Email   string `json:"email"`
	Tel     string `json:"tel"`
	ComID   int    `json:"com_id"`
}
