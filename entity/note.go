package entity

import "time"

type Note struct {
	Id          string     `json:"id"`
	DashboardId string     `json:"-" gorm:"not null"`
	OwnerUserId string     `json:"-" gorm:"not null"`
	OwnerUser   *User      `json:"ownerUser,omitempty" gorm:"foreignkey:OwnerUserId;constraint:onUpdate:CASCADE,onDelete:CASCADE"`
	Topic       string     `json:"topic" gorm:"not null"`
	Description string     `json:"description" gorm:"not null"`
	Status      bool       `json:"status" gorm:"not null"`
	Comments    *[]Comment `json:"comments,omitempty" gorm:"many2many:team_comments;constraint:onUpdate:CASCADE,onDelete:CASCADE"`
	CreatedAt   time.Time  `json:"createdAt"`
	UpdatedAt   time.Time  `json:"updatedAt"`
	DeadlineAt  time.Time  `json:"deadlineAt"`
}
