package repo

import (
	"../helper"
	"../model"
	"github.com/jinzhu/gorm"
)

// NoteRepo interface
type NoteRepo interface {
	Find(int) (*model.Note, error)
	List(helper.Pagination) ([]model.Note, error)
	Update(int, model.Note) error
	Delete(int) error
	Create(model.Note) (*model.Note, error)
}

// NoteRepoImpl struct
type NoteRepoImpl struct {
	DB *gorm.DB
}

// Create a note
func (noteRepo NoteRepoImpl) Create(note model.Note) (*model.Note, error) {
	err := noteRepo.DB.Create(&note).Error
	return &note, err
}

// Find a note
func (noteRepo NoteRepoImpl) Find(id int) (*model.Note, error) {
	note := &model.Note{}
	err := noteRepo.DB.Where("id = ?", id).First(note).Error
	return note, err
}

// List notes
func (noteRepo NoteRepoImpl) List(pagination helper.Pagination) ([]model.Note, error) {
	notes := []model.Note{}
	offset := pagination.GetOffset()
	limit := pagination.GetLimit()
	err := noteRepo.DB.Offset(offset).
		Limit(limit).
		Find(&notes).
		Error
	return notes, err
}

// Update a note
func (noteRepo NoteRepoImpl) Update(id int, note model.Note) error {
	err := noteRepo.DB.Where("id = ?", id).Update(&note).Error
	return err
}

// Delete a note
func (noteRepo NoteRepoImpl) Delete(id int) error {
	err := noteRepo.DB.Where("id = ?", id).Delete(&model.Note{}).Error
	return err
}
