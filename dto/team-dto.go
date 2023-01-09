package dto

type TeamCreateDTO struct {
	OwnerUserId string `json:"ownerUserId,omitempty"`
	Name        string `json:"name" binding:"required"`
}

type TeamAddMemberDTO struct {
	Email string `json:"email" binding:"required,email"`
}
