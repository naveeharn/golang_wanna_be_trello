package entity

import "time"

type User struct {
	Id          string    `json:"id" gorm:"primaryKey;type:varchar(24)"`
	FirstName   string    `json:"firstName" gorm:"type:varchar(64)"`
	LastName    string    `json:"lastName" gorm:"type:varchar(64)"`
	Email       string    `json:"email" gorm:"not null;uniqueIndex;type:varchar(128)"`
	Password    string    `json:"-" gorm:"->;<-;not null;type:varchar(128)"`
	BirthDate   time.Time `json:"birthDate"`
	Phone       string    `json:"phone" gorm:"type:varchar(12)"`
	CountryCode string    `json:"countryCode" gorm:"type:varchar(2)"`
	CreatedAt   time.Time `json:"createdAt"`
	UpdatedAt   time.Time `json:"updatedAt"`
	AccessToken string    `json:"accessToken,omitempty" gorm:"-"`
}
