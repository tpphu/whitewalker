package handler

import (
	"log"

	"github.com/tpphu/whitewalker/model"
	"github.com/tpphu/whitewalker/repo"
)

type noteHandlerImpl struct {
	noteRepo repo.NoteRepo
	log      *log.Logger
}

func (n noteHandlerImpl) get(id uint) (*model.Note, Error) {
	note, err := n.noteRepo.Find(id)
	if err != nil {
		return note, NewNotFoundErr(err)
	}
	return note, nil
}
