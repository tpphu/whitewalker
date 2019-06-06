package handler

import (
	"github.com/tpphu/whitewalker/model"
	"github.com/tpphu/whitewalker/repo"
)

type noteHandlerImpl struct {
	noteRepo repo.NoteRepo
}

func (n noteHandlerImpl) get(id int) (*model.Note, error) {
	return n.noteRepo.Find(id)
}
