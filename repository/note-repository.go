package repository

import "gorm.io/gorm"

type NoteRepository interface {
}

type noteConnection struct {
	connection *gorm.DB
}

func NewNoteRepository(db *gorm.DB) NoteRepository {
	return &noteConnection{
		connection: db,
	}
}
