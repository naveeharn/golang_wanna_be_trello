package entity

type Dashboard struct {
	Id          string `json:"id"`
	TeamId      string `json:"-" gorm:"not null;index:team_dashboard_name,unique"`
	Team        *Team  `json:"team,omitempty" gorm:"foreignkey:TeamId;constraint:onUpdate:CASCADE,onDelete:CASCADE"`
	Name        string `json:"name" gorm:"type:varchar(128);index:team_dashboard_name,unique"`
	OwnerUserId string `json:"-" gorm:"not null"`
	OwnerUser   *User  `json:"ownerUser,omitempty" gorm:"foreignkey:OwnerUserId;constraint:onUpdate:CASCADE,onDelete:CASCADE"`
	Notes       []Note `json:"notes,omitempty" gorm:"many2many:team_notes;constraint:onUpdate:CASCADE,onDelete:CASCADE"`
}
