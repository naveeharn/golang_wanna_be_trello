package repository

import (
	"fmt"

	"github.com/naveeharn/golang_wanna_be_trello/entity"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"gorm.io/gorm"
)

type TeamRepository interface {
	CreateTeam(team entity.Team) (entity.Team, error)
	GetTeamById(teamId, userId string) (entity.Team, error)
	GetTeamsByOwnerUserId(userId string) ([]entity.Team, error)
	AddMember(team entity.Team) (entity.Team, error)
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
	// team.Members = append(team.Members, entity.User{Id: team.OwnerUserId})
	transaction := db.connection.Begin()
	defer func() {
		if r := recover(); r != nil {
			transaction.Rollback()
		}
	}()
	if err := transaction.Error; err != nil {
		return entity.Team{}, err
	}

	db.connection.Create(&team)
	if transaction.Error != nil {
		return entity.Team{}, transaction.Error
	}
	if err := transaction.Association("Members").Append(&entity.User{Id: team.OwnerUserId}); err != nil {
		transaction.Rollback()
		return entity.Team{}, err
	}
	transaction = db.connection.Preload("OwnerUser").Preload("Members").Find(&team)
	if transaction.Error != nil {
		transaction.Rollback()
		return entity.Team{}, transaction.Error
	}
	return team, nil
}

func (db *teamConnection) GetTeamById(teamId, userId string) (entity.Team, error) {
	team := entity.Team{}
	transaction := db.connection.Raw("select team_id as id from team_members where team_id = ? and user_id = ?;", teamId, userId).Scan(&team)
	if team.Id == "" {
		return entity.Team{}, fmt.Errorf("user id cann't access to team data by team id")
	}
	transaction = transaction.Preload("OwnerUser").Preload("Members").Find(&team)
	if transaction.Error != nil {
		return entity.Team{}, transaction.Error
	}
	return team, nil
}

func (db *teamConnection) GetTeamsByOwnerUserId(ownerUserId string) ([]entity.Team, error) {
	teams := []entity.Team{}
	transaction := db.connection.Preload("OwnerUser").Preload("Members").Find(&teams, "owner_user_id = ?", ownerUserId)
	if transaction.Error != nil {
		return []entity.Team{}, transaction.Error
	}
	return teams, nil
}

func (db *teamConnection) AddMember(team entity.Team) (entity.Team, error) {

	transaction := db.connection.Save(&team)
	if transaction.Error != nil {
		return entity.Team{}, transaction.Error
	}

	return team, nil
}
