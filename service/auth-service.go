package service

import (
	"log"
	"runtime"

	"github.com/mashingan/smapping"
	"github.com/naveeharn/golang_wanna_be_trello/dto"
	"github.com/naveeharn/golang_wanna_be_trello/entity"
	"github.com/naveeharn/golang_wanna_be_trello/helper"
	"github.com/naveeharn/golang_wanna_be_trello/repository"
	"golang.org/x/crypto/bcrypt"
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
		helper.LoggerErrorPath(runtime.Caller(0))
		log.Fatalf("Failed to map struct: %v", err)
	}
	createdUser, err := service.userRepository.CreateUser(userBeforeCreate)
	return createdUser, err
}

func (service *authService) GetUserByEmail(email string) (entity.User, error) {
	user, err := service.userRepository.GetUserByEmail(email)
	return user, err
}

func (service *authService) IsDuplicateEmail(email string) bool {
	transaction := service.userRepository.IsDuplicateEmail(email)
	return transaction != nil
}

func (service *authService) VerifyCredential(email string, password string) interface{} {
	validatedUser, err := service.userRepository.GetUserByEmail(email)
	if err != nil {
		return nil
	}
	if !comparePassword(validatedUser.Password, password) {
		return nil
	}
	return validatedUser
}

func comparePassword(hashedPassword, plainPassword string) bool {
	if err := bcrypt.CompareHashAndPassword([]byte(hashedPassword), []byte(plainPassword)); err != nil {
		return false
	}
	return true
}
