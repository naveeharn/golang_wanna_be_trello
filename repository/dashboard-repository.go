package repository

import (
	"fmt"

	"github.com/naveeharn/golang_wanna_be_trello/entity"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"gorm.io/gorm"
)

type DashboardRepository interface {
	CreateDashboard(dashboard entity.Dashboard) (entity.Team, error)
}

type dashboardConnection struct {
	connection *gorm.DB
}

func NewDashboardRepository(db *gorm.DB) DashboardRepository {
	return &dashboardConnection{
		connection: db,
	}
}

func (db *dashboardConnection) CreateDashboard(dashboard entity.Dashboard) (entity.Team, error) {
	team := entity.Team{}
	dashboard.Id = primitive.NewObjectID().Hex()
	err := db.connection.Transaction(func(transaction *gorm.DB) error {
		user := entity.User{Id: dashboard.OwnerUserId}
		transaction.Where(&user).First(&user)
		if user.Id == "" {
			return fmt.Errorf("user id doesn't found")
		}
		transaction.Where(&team).First(&team, "owner_user_id = ?", dashboard.OwnerUserId)
		if team.OwnerUserId == "" {
			return fmt.Errorf("user id doesn't allow to create new dashboard to team")
		}
		transaction.Create(&dashboard)
		if transaction.Error != nil {
			return transaction.Error
		}
		transaction.Model(&team).Association("Dashboards").Append(&dashboard)
		if transaction.Error != nil {
			return transaction.Error
		}
		transaction.Preload("OwnerUser").Preload("Members").Preload("Dashboards.OwnerUser").Preload("Dashboards.Team.OwnerUser").Find(&team)
		if transaction.Error != nil {
			return transaction.Error
		}
		return nil
	})
	if err != nil {
		return entity.Team{}, err
	}
	return team, nil
}
