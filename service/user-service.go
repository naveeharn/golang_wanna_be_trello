package service

import (
	"fmt"

	"github.com/mashingan/smapping"
	"github.com/naveeharn/golang_wanna_be_trello/dto"
	"github.com/naveeharn/golang_wanna_be_trello/entity"
	"github.com/naveeharn/golang_wanna_be_trello/repository"
)

type UserService interface {
	Update(user dto.UserUpdateDTO) (entity.User, error)
	GetUserById(userId string) (entity.User, error)
	ResetPassword(user dto.UserResetPasswordDTO) (entity.User, error)
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
	user, err := service.userRepository.GetUserById(userId)
	return user, err
}

func (service *userService) ResetPassword(user dto.UserResetPasswordDTO) (entity.User, error) {
	if user.Id == "" {
		return entity.User{}, fmt.Errorf("User.Id doesn't exists")
	}

	userBeforeUpdate, err := service.userRepository.GetUserById(user.Id)
	if err != nil {
		return entity.User{}, err
	}
	if !comparePassword(userBeforeUpdate.Password, user.OldPassword) {
		return entity.User{}, fmt.Errorf("password of user is not correct")
	}
	userBeforeUpdate.Password = user.NewPassword
	updatedUser, err := service.userRepository.ResetPassword(userBeforeUpdate)
	if err != nil {
		return entity.User{}, err
	}
	return updatedUser, err
}

func (service *userService) Update(user dto.UserUpdateDTO) (entity.User, error) {
	if user.Id == "" {
		return entity.User{}, fmt.Errorf("User.Id doesn't exists")
	}

	userBeforeUpdate, err := service.userRepository.GetUserById(user.Id)
	if err != nil {
		return entity.User{}, err
	}

	err = smapping.FillStruct(&userBeforeUpdate, smapping.MapFields(&user))
	if err != nil {
		return entity.User{}, err
	}

	updatedUser, err := service.userRepository.UpdateUser(userBeforeUpdate)
	if err != nil {
		return entity.User{}, err
	}

	return updatedUser, nil
}
