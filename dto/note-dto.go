package dto

import "time"

type NoteCreateDTO struct {
	DashboardId string `json:"dashboardId,omitempty"`
	OwnerUserId string `json:"ownerUserId,omitempty"`
	Topic       string `json:"topic" binding:"required,max=128"`
	Description string `json:"description" binding:"required,max=1024"`
	// Status      *bool     `json:"status" binding:"required"`
	DeadlineAt time.Time `json:"deadlineAt" binding:"required"`
}

type NoteUpdateDTO struct {
	Id          string    `json:"-"`
	DashboardId string    `json:"dashboardId,omitempty"`
	OwnerUserId string    `json:"ownerUserId,omitempty"`
	Topic       string    `json:"topic" binding:"required,max=128"`
	Description string    `json:"description" binding:"required,max=1024"`
	Status      *bool     `json:"status" binding:"required"`
	DeadlineAt  time.Time `json:"deadlineAt" binding:"required"`
}

type NoteDeleteDTO struct {
	Id          string `json:"-"`
	DashboardId string `json:"dashboardId,omitempty"`
	OwnerUserId string `json:"ownerUserId,omitempty"`
}
