package entity

import "time"

type User struct {
	Id          string    `json:"id" gorm:"primaryKey"`
	FirstName   string    `json:"firstName" gorm:"type:varchar(64)"`
	LastName    string    `json:"lastName" gorm:"type:varchar(64)"`
	Email       string    `json:"email" gorm:"not null;uniqueIndex;type:varchar(128)"`
	Password    string    `json:"-" gorm:"->;<-;not null"`
	BirthDate   time.Time `json:"birthDate"`
	Phone       string    `json:"phone" binding:"required,e164"`
	CountryCode string    `json:"countryCode" binding:"required,iso3166_1_alpha2"`
	CreatedAt   time.Time `json:"createdAt"`
	UpdatedAt   time.Time `json:"updatedAt"`
	AccessToken string    `json:"accessToken,omitempty" gorm:"-"`
}
