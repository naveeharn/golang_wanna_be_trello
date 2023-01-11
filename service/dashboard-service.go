package service

import (
	"fmt"

	"github.com/mashingan/smapping"
	"github.com/naveeharn/golang_wanna_be_trello/dto"
	"github.com/naveeharn/golang_wanna_be_trello/entity"
	"github.com/naveeharn/golang_wanna_be_trello/repository"
)

type DashboardService interface {
	CreateDashboard(dashboard dto.DashboardCreateDTO) (entity.Team, error)
}

type dashboardService struct {
	dashboardRepository repository.DashboardRepository
}

func NewDashboardService(dashboardRepository repository.DashboardRepository) DashboardService {
	return &dashboardService{
		dashboardRepository: dashboardRepository,
	}
}

func (service *dashboardService) CreateDashboard(dashboard dto.DashboardCreateDTO) (entity.Team, error) {
	if dashboard.OwnerUserId == "" {
		return entity.Team{}, fmt.Errorf("dashboard.OwnerUserId doesn't exists")
	}
	dashboardBeforeCreate := entity.Dashboard{}
	err := smapping.FillStruct(&dashboardBeforeCreate, smapping.MapFields(&dashboard))
	if err != nil {
		return entity.Team{}, err
	}
	updatedTeam, err := service.dashboardRepository.CreateDashboard(dashboardBeforeCreate)
	if err != nil {
		return entity.Team{}, err
	}
	return updatedTeam, nil
}
