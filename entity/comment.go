package entity

import "time"

type Comment struct {
	Id          string    `json:"id"`
	TeamId      string    `json:"-" gorm:"not null"`
	NoteId      string    `json:"noteId" gorm:"not null"`
	OwnerUserId string    `json:"-" gorm:"not null"`
	OwnerUser   User      `json:"ownerUser,omitempty" gorm:"foreignkey:OwnerUserId;constraint:onUpdate:CASCADE,onDelete:CASCADE"`
	Description string    `json:"description" gorm:"not null"`
	CreatedAt   time.Time `json:"createdAt"`
	UpdatedAt   time.Time `json:"updatedAt"`
}
