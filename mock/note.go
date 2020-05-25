package mock

import (
	"github.com/stretchr/testify/mock"
	"github.com/tpphu/whitewalker/helper"
	"github.com/tpphu/whitewalker/model"
)

type NoteRepoImpl struct {
	mock.Mock
}

func (self *NoteRepoImpl) Create(note model.Note) (*model.Note, error) {
	args := self.Called(note)
	out := args.Get(0)
	if out == nil {
		return nil, args.Error(1)
	}
	return out.(*model.Note), args.Error(1)
}

func (self *NoteRepoImpl) Find(id uint) (*model.Note, error) {
	args := self.Called(id)
	out := args.Get(0)
	if out == nil {
		return nil, args.Error(1)
	}
	return out.(*model.Note), args.Error(1)
}

func (self *NoteRepoImpl) List(pagination helper.Pagination) ([]model.Note, error) {
	args := self.Called(pagination)
	out := args.Get(0)
	if out == nil {
		return nil, args.Error(1)
	}
	return out.([]model.Note), args.Error(1)
}

func (self *NoteRepoImpl) Update(id uint, note model.Note) error {
	args := self.Called(id, note)
	return args.Error(0)
}

func (self *NoteRepoImpl) Delete(id uint) error {
	args := self.Called(id)
	return args.Error(0)
}
