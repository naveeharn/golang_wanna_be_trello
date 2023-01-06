package dto

import (
	"time"
)

type UserCreateDTO struct {
	FirstName   string    `json:"firstName" binding:"required,alpha" validate:"min:1;max:128"`
	LastName    string    `json:"lastName" binding:"required,alpha" validate:"min:1;max:128"`
	Email       string    `json:"email" binding:"required,email"`
	Password    string    `json:"password" binding:"required" validate:"min:8"`
	BirthDate   time.Time `json:"birthDate" binding:"required" time_format:"2006-02-01"`
	Phone       string    `json:"phone" binding:"required,e164"`
	CountryCode string    `json:"countryCode" binding:"required,iso3166_1_alpha2"`
}

type UserUpdateDTO struct {
	Id          string    `json:"-"`
	FirstName   string    `json:"firstName" binding:"required,alpha" validate:"min:1;max:128"`
	LastName    string    `json:"lastName" binding:"required,alpha" validate:"min:1;max:128"`
	BirthDate   time.Time `json:"birthDate" binding:"required" time_format:"2006-02-01"`
	Phone       string    `json:"phone" binding:"required,e164"`
	CountryCode string    `json:"countryCode" binding:"required,iso3166_1_alpha2"`
}

type UserResetPasswordDTO struct {
	Id          string `json:"-"`
	OldPassword string `json:"oldPassword" binding:"required"`
	NewPassword string `json:"newPassword" binding:"required,min=3"`
}
