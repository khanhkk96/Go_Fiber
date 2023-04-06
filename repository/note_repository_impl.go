package repository

import (
	"errors"
	"go-fiber-crud/data/request"
	"go-fiber-crud/helper"
	"go-fiber-crud/model"

	"gorm.io/gorm"
)

type NoteRepositoryImpl struct {
	Db *gorm.DB
}

// Delete implements NoteRepository
func (n *NoteRepositoryImpl) Delete(noteId int) {
	//panic("unimplemented")
	var note model.Note
	result := n.Db.Where("id = ?", noteId).Delete(&note)
	helper.ErrorPanic(result.Error)
}

// FindAll implements NoteRepository
func (n *NoteRepositoryImpl) FindAll() []model.Note {
	//panic("unimplemented")
	var note []model.Note
	result := n.Db.Find(&note)
	helper.ErrorPanic(result.Error)
	return note
}

// FindById implements NoteRepository
func (n *NoteRepositoryImpl) FindById(noteId int) (model.Note, error) {
	//panic("unimplemented")
	var note model.Note
	result := n.Db.Find(&note, noteId)
	if result != nil {
		return note, nil
	}
	return note, errors.New("Note is not found")
}

// Save implements NoteRepository
func (n *NoteRepositoryImpl) Save(note model.Note) {
	//panic("unimplemented")
	result := n.Db.Create(&note)
	helper.ErrorPanic(result.Error)
}

// Update implements NoteRepository
func (n *NoteRepositoryImpl) Update(note model.Note) {
	//panic("unimplemented")
	var updateNote = request.UpdateNoteRequest{
		Id:      note.Id,
		Content: note.Content,
	}
	result := n.Db.Model(&note).Updates(updateNote)
	helper.ErrorPanic(result.Error)
}

func NewNoteRepositoryImpl(Db *gorm.DB) NoteRepository {
	return &NoteRepositoryImpl{Db: Db}
}
