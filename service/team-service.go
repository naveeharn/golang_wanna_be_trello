package service

import (
	"fmt"

	"github.com/mashingan/smapping"
	"github.com/naveeharn/golang_wanna_be_trello/dto"
	"github.com/naveeharn/golang_wanna_be_trello/entity"
	"github.com/naveeharn/golang_wanna_be_trello/repository"
)

type TeamService interface {
	CreateTeam(team dto.TeamCreateDTO) (entity.Team, error)
	GetTeamById(teamId, userId string) (entity.Team, error)
	GetTeamsByOwnerUserId(ownerUserId string) ([]entity.Team, error)
	AddMember(team entity.Team) (entity.Team, error)
}

type teamService struct {
	teamRepository repository.TeamRepository
}

func NewTeamService(teamRepository repository.TeamRepository) TeamService {
	return &teamService{
		teamRepository: teamRepository,
	}
}

func (service *teamService) AddMember(team entity.Team) (entity.Team, error) {
	panic("unimplemented")
}

func (service *teamService) CreateTeam(team dto.TeamCreateDTO) (entity.Team, error) {
	if team.OwnerUserId == "" {
		return entity.Team{}, fmt.Errorf("Team.OwnerUserId doesn't exists")
	}
	teamBeforeCreate := entity.Team{}
	err := smapping.FillStruct(&teamBeforeCreate, smapping.MapFields(&team))
	if err != nil {
		return entity.Team{}, err
	}
	createdTeam, err := service.teamRepository.CreateTeam(teamBeforeCreate)
	if err != nil {
		return entity.Team{}, err
	}
	return createdTeam, nil
}

func (service *teamService) GetTeamById(teamId, userId string) (entity.Team, error) {
	team, err := service.teamRepository.GetTeamById(teamId, userId)
	if err != nil {
		return entity.Team{}, err
	}
	return team, nil
}

func (service *teamService) GetTeamsByOwnerUserId(ownerUserId string) ([]entity.Team, error) {
	teams, err := service.teamRepository.GetTeamsByOwnerUserId(ownerUserId)
	if err != nil {
		return []entity.Team{}, err
	}
	return teams, nil
}
