package models

import "gorm.io/gorm"

type user struct{
	gorm.Model
	Id uint `json:"id"`
	FirstName string `json:"first_name"`
	LastName string `json:"last_name"`
	Email string `json:"email"`
}