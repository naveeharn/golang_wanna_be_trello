package entity

type Dashboard struct {
	Id          string `json:"id"`
	TeamId      string `json:"-" gorm:"not null;index:team_dashboard_name,unique"`
	Team        Team   `json:"team" gorm:"foreignkey:TeamId;constraint:onUpdate:CASCADE,onDelete:CASCADE"`
	Name        string `json:"name" gorm:"type:varchar(128);index:team_dashboard_name,unique"`
	OwnerUserId string `json:"-" gorm:"not null"`
	OwnerUser   User   `json:"ownerUser" gorm:"foreignkey:OwnerUserId;constraint:onUpdate:CASCADE,onDelete:CASCADE"`
	// notes *[]Note
}
