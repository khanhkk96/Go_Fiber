package repository

import "go-fiber-crud/model"

type NoteRepository interface {
	Save(note model.Note)
	Update(note model.Note)
	Remove(noteId int)
	Delete(noteId int)
	FindById(noteId int) (model.Note, error)
	FindAll() []model.Note
}
