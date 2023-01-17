package repository

import (
	"fmt"
	"log"
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
		user := entity.User{Id: note.OwnerUserId}
		if err := transaction.Model(&team).Association("Members").Find(&user); err != nil || user == (entity.User{Id: note.OwnerUserId}) {
			return fmt.Errorf("User.Id can't allow to create note in dashboard")
		}
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
	err := db.connection.Transaction(func(transaction *gorm.DB) error {
		noteBeforeUpdate := entity.Note{Id: note.Id}
		transaction.Find(&noteBeforeUpdate)
		if transaction.Error != nil {
			return transaction.Error
		}

		dashboard.Id = noteBeforeUpdate.DashboardId
		transaction.Find(&dashboard)
		if transaction.Error != nil {
			return transaction.Error
		}

		team := entity.Team{Id: dashboard.TeamId}
		user := entity.User{Id: note.OwnerUserId}
		if err := transaction.Model(&team).Association("Members").Find(&user); err != nil || user == (entity.User{Id: note.OwnerUserId}) {
			return fmt.Errorf("User.Id can't allow to create note in dashboard")
		}
		// log.Printf("\n\n%#v\n\n", team)
		// log.Printf("\n\n%#v\n\n", user)
		// if err := transaction.Model(&team).Association("Members").Find(&user); err != nil || user == (entity.User{Id: note.OwnerUserId}) {
		// 	log.Printf("\n\n%#v\n\n", user)
		// 	log.Printf("\n\n%#v\n\n", user == (entity.User{Id: note.OwnerUserId}))
		// 	return fmt.Errorf("User.Id can't allow to create note in dashboard")
		// }
		// log.Printf("\n\n%#v\n\n", user)
		// log.Printf("\n\n%#v\n\n", user == (entity.User{Id: note.OwnerUserId}))
		note.CreatedAt = noteBeforeUpdate.CreatedAt
		transaction.Save(&note)
		if transaction.Error != nil {
			return transaction.Error
		}

		transaction.Preload("OwnerUser").Preload("Notes").Find(&dashboard)
		if transaction.Error != nil {
			return transaction.Error
		}

		return nil
	})
	if err != nil {
		return entity.Dashboard{}, err
	}
	return dashboard, nil
}

func (db *noteConnection) DeleteNote(note entity.Note) (entity.Dashboard, error) {
	// note := entity.Note{Id: noteId}
	dashboard := entity.Dashboard{}

	err := db.connection.Transaction(func(transaction *gorm.DB) error {
		noteBeforeDelete := entity.Note{Id: note.Id}
		transaction.Find(&noteBeforeDelete)
		if transaction.Error != nil {
			return transaction.Error
		}
		// log.Println(note.Id
		log.Printf("\n\n%#v\n\n", noteBeforeDelete)
		log.Println(">>>", note.DashboardId)
		log.Println(">>>", noteBeforeDelete.DashboardId)

		if note.DashboardId != noteBeforeDelete.DashboardId {
			return fmt.Errorf("dashboardId of note doesn't exists")
		}

		log.Println(note.Id)
		dashboard.Id = noteBeforeDelete.DashboardId
		transaction.Find(&dashboard)
		if transaction.Error != nil {
			return transaction.Error
		}

		team := entity.Team{Id: dashboard.TeamId}
		user := entity.User{Id: note.OwnerUserId}
		log.Println(note.Id)
		if err := transaction.Model(&team).Association("Members").Find(&user); err != nil || user == (entity.User{Id: note.OwnerUserId}) {
			return fmt.Errorf("User.Id can't allow to create note in dashboard")
		}
		// if err != transaction.Model(&dashboard).Association("Notes") {
		// 	return nil, err
		// }
		log.Println(note.Id)
		transaction.Delete(&note, "id = ?", note.Id)
		if transaction.Error != nil {
			return transaction.Error
		}

		transaction.Preload("OwnerUser").Preload("Notes").Find(&dashboard)
		if transaction.Error != nil {
			return transaction.Error
		}
		return nil
	})

	if err != nil {
		return entity.Dashboard{}, err
	}
	return dashboard, nil
}
