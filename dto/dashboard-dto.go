package dto

type DashboardCreateDTO struct {
	TeamId      string `json:"teamId,omitempty"`
	OwnerUserId string `json:"ownerUserId,omitempty"`
	Name        string `json:"name" binding:"required"`
}
