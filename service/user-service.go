package service

import (
	"github.com/naveeharn/golang_wanna_be_trello/dto"
	"github.com/naveeharn/golang_wanna_be_trello/entity"
	"github.com/naveeharn/golang_wanna_be_trello/repository"
)

type UserService interface {
	Update(user dto.UserUpdateDTO) (entity.User, error)
	GetUserById(userId string) (entity.User, error)
}

type userService struct {
	userRepository repository.UserRepository
}

func NewUserService(userRepository repository.UserRepository) UserService {
	return &userService{
		userRepository: userRepository,
	}
}

func (service *userService) GetUserById(userId string) (entity.User, error) {
	panic("unimplemented")
}

func (service *userService) Update(user dto.UserUpdateDTO) (entity.User, error) {
	panic("unimplemented")
}
