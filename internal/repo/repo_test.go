package repo

import (
	"testing"
	"time"

	"github.com/stretchr/testify/suite"

	"github.com/akhilesharora/todo/internal/utils"
)

type Suite struct {
	suite.Suite
	repo *TodoService
}

func (s *Suite) Test01CreateNote() {
	note := &Note{
		ID:       1,
		Title:    "Hello World",
		Comments: "The earth is round",
		DueDate:  time.Now(),
	}
	err := s.repo.Create(note)
	s.Assert().NoError(err)
}

func (s *Suite) Test02GetAllNotes() {
	s.repo.notes = []*Note{
		{
			ID:       1,
			Title:    "Hello World",
			Comments: "Check if Earth is like egg shape?",
			DueDate:  time.Now(),
		}, {
			ID:       2,
			Title:    "Mars",
			Comments: "Check life on Mars?",
			DueDate:  time.Now(),
		},
	}

	notes, err := s.repo.GetAll()
	s.Assert().NoError(err)
	s.Assert().Equal(2, len(notes))
}

func (s *Suite) Test03CreateEmptyNote() {
	err := s.repo.Create(nil)
	s.Assert().EqualError(err, utils.ErrEmptyNote.Error())
}

func (s *Suite) Test04UpdateNote() {
	note := &Note{
		ID:       2,
		Title:    "Venus",
		Comments: "Is venus too hot?",
		DueDate:  time.Now(),
	}
	updateNote, err := s.repo.Update(note)
	s.Assert().NoError(err)
	s.Assert().Equal(note.ID, updateNote.ID)
	s.Assert().Equal(note.Title, updateNote.Title)
	s.Assert().Equal(note.Comments, updateNote.Comments)
}

func (s *Suite) Test05UpdateNoteWithInvalidId() {
	note := Note{
		ID:       5,
		Title:    "Venus",
		Comments: "Is venus too hot?",
		DueDate:  time.Now(),
	}
	newNote, err := s.repo.Update(&note)
	s.Assert().Error(err)
	s.Assert().Nil(newNote)
	s.Assert().EqualError(err, utils.ErrNotFound.Error())
}

func (s *Suite) Test06DeleteNote() {
	notes, err := s.repo.GetAll()
	s.Assert().NoError(err)
	noteId := notes[0].ID
	err = s.repo.Delete(noteId)
	s.Assert().NoError(err)
}

func (s *Suite) Test07DeleteNoteWithInvalidId() {
	noteId := uint32(len(s.repo.notes) + 1000)
	err := s.repo.Delete(noteId)
	s.Assert().Error(err)
	s.Assert().EqualError(err, utils.ErrNotFound.Error())
}

func (s *Suite) Test08GetNote() {
	notes, err := s.repo.GetAll()
	validNoteId := notes[0].ID
	note, err := s.repo.get(validNoteId)
	s.Assert().NoError(err)
	s.Assert().Equal(note.ID, validNoteId)
}

func (s *Suite) Test09GetAllNoteWhenEmpty() {
	s.repo.notes = []*Note{}
	notes, err := s.repo.GetAll()
	s.Assert().Equal(0, len(notes))
	s.Assert().EqualError(err, utils.ErrEmptyNotes.Error())
}

func TestSuite(t *testing.T) {
	testSuite := new(Suite)
	testSuite.repo = NewTodoService()
	suite.Run(t, testSuite)
}
