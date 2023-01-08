package repository

import (
	"runtime"

	"github.com/naveeharn/golang_wanna_be_trello/entity"
	"github.com/naveeharn/golang_wanna_be_trello/helper"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"gorm.io/gorm"
)

type TeamRepository interface {
	CreateTeam(team entity.Team) (entity.Team, error)
}

type teamConnection struct {
	connection *gorm.DB
}

func NewTeamRepository(db *gorm.DB) TeamRepository {
	return &teamConnection{
		connection: db,
	}
}

func (db *teamConnection) CreateTeam(team entity.Team) (entity.Team, error) {
	team.Id = primitive.NewObjectID().Hex()
	helper.LoggerErrorPath(runtime.Caller(0))
	// team.Members = append((team.Members).([]entity.User), entity.User{Id: team.OwnerUserId})
	transaction := db.connection.Save(&team)
	if transaction.Error != nil {
		return entity.Team{}, transaction.Error
	}
	team.Members = append(team.Members, entity.User{Id: team.OwnerUserId})

	transaction = db.connection.Preload("OwnerUser").Preload("Members").Find(&team)
	if transaction.Error != nil {
		return entity.Team{}, transaction.Error
	}
	return team, nil
}
