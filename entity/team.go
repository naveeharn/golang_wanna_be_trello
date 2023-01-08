package entity

import "time"

type Team struct {
	Id          string    `json:"id"`
	OwnerUserId string    `json:"-" gorm:"not null"`
	OwnerUser   User      `json:"ownerUser" gorm:"foreignkey:OwnerUserId;constraint:onUpdate:CASCADE,onDelete:CASCADE"`
	Name        string    `json:"name"`
	Members     []User    `json:"members,omitempty" gorm:"many2many:team_members"`
	CreatedAt   time.Time `json:"createdAt"`
	UpdatedAt   time.Time `json:"updatedAt"`
}
