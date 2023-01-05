package service

import (
	"log"

	"github.com/mashingan/smapping"
	"github.com/naveeharn/golang_wanna_be_trello/dto"
	"github.com/naveeharn/golang_wanna_be_trello/entity"
	"github.com/naveeharn/golang_wanna_be_trello/repository"
)

type AuthService interface {
	VerifyCredential(email, password string) interface{}
	CreateUser(user dto.UserCreateDTO) (entity.User, error)
	GetUserByEmail(email string) (entity.User, error)
	IsDuplicateEmail(email string) bool
}

type authService struct {
	userRepository repository.UserRepository
}

func NewAuthService(userRepository repository.UserRepository) AuthService {
	return &authService{
		userRepository: userRepository,
	}
}

func (service *authService) CreateUser(user dto.UserCreateDTO) (entity.User, error) {
	userBeforeCreate := entity.User{}
	if err := smapping.FillStruct(&userBeforeCreate, smapping.MapFields(&user)); err != nil {
		log.Fatalf("Failed to map struct: %v", err)
	}
	createdUser, err := service.userRepository.CreateUser(userBeforeCreate)
	return createdUser, err
}

func (service *authService) GetUserByEmail(email string) (entity.User, error) {
	panic("unimplemented")
}

func (service *authService) IsDuplicateEmail(email string) bool {
	transaction := service.userRepository.IsDuplicateEmail(email)
	return transaction != nil
}

func (service *authService) VerifyCredential(email string, password string) interface{} {
	panic("unimplemented")
}
