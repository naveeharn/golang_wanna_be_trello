package repository

import (
	"fmt"
	"time"

	"github.com/naveeharn/golang_wanna_be_trello/entity"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"gorm.io/gorm"
)

type NoteRepository interface {
	CreateNote(note entity.Note) (entity.Dashboard, error)
	UpdateNote(note entity.Note) (entity.Dashboard, error)
	DeleteNote(note entity.Note) (entity.Dashboard, error)
}

type noteConnection struct {
	connection *gorm.DB
}

func NewNoteRepository(db *gorm.DB) NoteRepository {
	return &noteConnection{
		connection: db,
	}
}

func (db *noteConnection) CreateNote(note entity.Note) (entity.Dashboard, error) {
	dashboard := entity.Dashboard{Id: note.DashboardId}
	err := db.connection.Transaction(func(transaction *gorm.DB) error {
		transaction.Find(&dashboard)
		if transaction.Error != nil {
			return transaction.Error
		}
		team := entity.Team{Id: dashboard.TeamId}
		if err := transaction.Model(&team).Association("Members").Find(&entity.User{Id: note.OwnerUserId}); err != nil {
			return err
		}
		// if transaction.Error != nil {
		// 	return transaction.Error
		// }
		note.Id = primitive.NewObjectID().Hex()
		transaction.Create(&note)
		if transaction.Error != nil {
			return transaction.Error
		}
		if err := transaction.Model(&dashboard).Association("Notes").Append(&note); err != nil {
			return err
		}
		if note.DeadlineAt.Before(time.Now()) {
			return fmt.Errorf("deadline can't before now")
		}
		transaction.Preload("OwnerUser").Preload("Notes").Find(&dashboard)
		return nil
	})
	if err != nil {
		return entity.Dashboard{}, err
	}

	return dashboard, nil
}

func (db *noteConnection) UpdateNote(note entity.Note) (entity.Dashboard, error) {
	dashboard := entity.Dashboard{}
	return dashboard, nil
}

func (db *noteConnection) DeleteNote(note entity.Note) (entity.Dashboard, error) {
	dashboard := entity.Dashboard{}
	return dashboard, nil
}
