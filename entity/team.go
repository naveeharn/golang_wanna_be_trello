package entity

import "time"

type Team struct {
	Id          string      `json:"id" gorm:"primaryKey;type:varchar(24)"`
	OwnerUserId string      `json:"-" gorm:"not null;index:owner_team_name,unique"`
	OwnerUser   *User       `json:"ownerUser" gorm:"foreignkey:OwnerUserId;constraint:onUpdate:CASCADE,onDelete:CASCADE"`
	Name        string      `json:"name" gorm:"type:varchar(128);index:owner_team_name,unique"`
	Members     []User      `json:"members,omitempty" gorm:"many2many:team_members;constraint:onUpdate:CASCADE,onDelete:CASCADE"`
	Dashboards  []Dashboard `json:"dashboards,omitempty" gorm:"many2many:team_dashboards;constraint:onUpdate:CASCADE,onDelete:CASCADE"`
	CreatedAt   time.Time   `json:"createdAt"`
	UpdatedAt   time.Time   `json:"updatedAt"`
}
