package repository

import (
	"fmt"
	"runtime"

	"github.com/naveeharn/golang_wanna_be_trello/entity"
	"github.com/naveeharn/golang_wanna_be_trello/helper"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"gorm.io/gorm"
)

type DashboardRepository interface {
	CreateDashboard(dashboard entity.Dashboard) (entity.Team, error)
	UpdateDashboard(dashboard entity.Dashboard) (entity.Team, error)
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
		// team = entity.Team{Id: team.Id}
		transaction.Preload("OwnerUser").Preload("Members").Preload("Dashboards"). /*Preload("Dashboards.OwnerUser").Preload("Dashboards.Team.OwnerUser")*/ Find(&team)
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

func (db *dashboardConnection) UpdateDashboard(dashboard entity.Dashboard) (entity.Team, error) {
	team := entity.Team{}
	err := db.connection.Transaction(func(transaction *gorm.DB) error {
		user := entity.User{Id: dashboard.OwnerUserId}
		transaction.Where(&user).Find(&user)
		if transaction.Error != nil {
			helper.LoggerErrorPath(runtime.Caller(0))
			return transaction.Error
		}
		dashboardBeforeUpdate := entity.Dashboard{Id: dashboard.Id}
		transaction.Find(&dashboardBeforeUpdate, "id = ? and owner_user_id = ?", dashboard.Id, user.Id)
		if transaction.Error != nil {
			helper.LoggerErrorPath(runtime.Caller(0))
			return transaction.Error
		}
		dashboardBeforeUpdate.Name = dashboard.Name
		transaction.Save(&dashboardBeforeUpdate)
		if transaction.Error != nil {
			helper.LoggerErrorPath(runtime.Caller(0))
			return transaction.Error
		}
		transaction.Preload("OwnerUser").Preload("Members").Preload("Dashboards").Find(&team)
		// log.Printf("%#v\n\n", team)
		if transaction.Error != nil {
			helper.LoggerErrorPath(runtime.Caller(0))
			return transaction.Error
		}
		return nil
	})
	if err != nil {
		return entity.Team{}, err
	}
	return team, nil
}
