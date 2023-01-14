package service

import (
	"fmt"

	"github.com/mashingan/smapping"
	"github.com/naveeharn/golang_wanna_be_trello/dto"
	"github.com/naveeharn/golang_wanna_be_trello/entity"
	"github.com/naveeharn/golang_wanna_be_trello/repository"
)

type NoteService interface {
	CreateNote(note dto.NoteCreateDTO) (entity.Dashboard, error)
	UpdateNote(note dto.NoteUpdateDTO) (entity.Dashboard, error)
	DeleteNote(note dto.NoteDeleteDTO) (entity.Dashboard, error)
}

type noteService struct {
	noteRepository repository.NoteRepository
}

func NewNoteService(noteRepository repository.NoteRepository) NoteService {
	return &noteService{
		noteRepository: noteRepository,
	}
}

func (service *noteService) CreateNote(note dto.NoteCreateDTO) (entity.Dashboard, error) {
	if note.OwnerUserId == "" {
		return entity.Dashboard{}, fmt.Errorf("Note.OwnerUserId doesn't exists")
	}
	noteBeforeCreate := entity.Note{}
	err := smapping.FillStruct(&noteBeforeCreate, smapping.MapFields(&note))
	if err != nil {
		return entity.Dashboard{}, err
	}
	updatedDashboard, err := service.noteRepository.CreateNote(noteBeforeCreate)
	if err != nil {
		return entity.Dashboard{}, err
	}
	return updatedDashboard, nil
}

func (service *noteService) DeleteNote(note dto.NoteDeleteDTO) (entity.Dashboard, error) {
	return entity.Dashboard{}, nil
}

func (service *noteService) UpdateNote(note dto.NoteUpdateDTO) (entity.Dashboard, error) {
	return entity.Dashboard{}, nil
}
