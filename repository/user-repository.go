package repository

import (
	"fmt"
	"log"
	"runtime"

	"github.com/naveeharn/golang_wanna_be_trello/entity"
	"github.com/naveeharn/golang_wanna_be_trello/helper"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

type UserRepository interface {
	CreateUser(user entity.User) (entity.User, error)
	UpdateUser(user entity.User) (entity.User, error)
	VerifyCredential(email, password string) interface{}
	IsDuplicateEmail(email string) (tx *gorm.DB)
	GetUserByEmail(email string) (entity.User, error)
	GetUserById(userId string) (entity.User, error)
	ResetPassword(user entity.User) (entity.User, error)
}

type userConnection struct {
	connection *gorm.DB
}

func NewUserRepository(db *gorm.DB) UserRepository {
	return &userConnection{
		connection: db,
	}
}

func (db *userConnection) CreateUser(user entity.User) (entity.User, error) {
	user.Id = primitive.NewObjectID().Hex()
	user.Password = hashAndSalt(user.Password)
	transaction := db.connection.Create(&user)
	if transaction.Error != nil {
		return entity.User{}, transaction.Error
	}
	return user, nil
}

func (db *userConnection) GetUserByEmail(email string) (entity.User, error) {
	user := entity.User{}
	transaction := db.connection.Where("email = ?", email).Take(&user)
	if transaction.Error != nil {
		return entity.User{}, transaction.Error
	}
	return user, nil
}

func (db *userConnection) GetUserById(userId string) (entity.User, error) {
	user := entity.User{}

	transaction := db.connection.Find(&user, "id = ?", userId)
	if transaction.Error != nil {
		return entity.User{}, transaction.Error
	}
	if user == (entity.User{}) {
		return entity.User{}, fmt.Errorf("user by id not found")
	}
	return user, nil
}

func (db *userConnection) IsDuplicateEmail(email string) (tx *gorm.DB) {
	user := entity.User{}
	transaction := db.connection.Where("email = ?", email).Take(&user)
	if transaction.Error != nil {
		return nil
	}
	return transaction
}

func (db *userConnection) UpdateUser(user entity.User) (entity.User, error) {
	transaction := db.connection.Save(&user)
	if transaction.Error != nil {
		return entity.User{}, transaction.Error
	}
	return user, nil
}

func (db *userConnection) ResetPassword(user entity.User) (entity.User, error) {
	user.Password = hashAndSalt(user.Password)
	return db.UpdateUser(user)
}

func (db *userConnection) VerifyCredential(email, password string) interface{} {
	user := entity.User{}
	transaction := db.connection.Where("email = ?", email).Take(&user)
	if transaction.Error != nil {
		return nil
	}
	panic("unimplemented")
}

func hashAndSalt(password string) string {
	hashed, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		helper.LoggerErrorPath(runtime.Caller(0))
		log.Fatalf("Failed to hash a password %s\n error:%s", password, err.Error())
	}
	return string(hashed)
}
