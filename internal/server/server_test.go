package server

import (
	"context"
	"testing"

	"github.com/pkg/errors"
	"github.com/stretchr/testify/suite"

	"github.com/akhilesharora/todo/internal/repo"
	"github.com/akhilesharora/todo/internal/utils"
	"github.com/akhilesharora/todo/pb"
)

var ctx = context.Background()

type Suite struct {
	suite.Suite
	server *Server
}

func (s *Suite) Test01CreateNote() {
	note := &pb.Note{
		Title:    "Meteor Shower",
		Comments: "checkout Nasa.gov",
		DueDate:  "2019-10-21",
	}
	res, err := s.server.CreateNote(ctx, &pb.CreateNoteMsg{
		Note: note,
	})
	s.Assert().NoError(err)
	s.Assert().Equal(&pb.Empty{}, res)
}

func (s *Suite) Test02CreateNoteWithEmptyTitle() {
	note := &pb.Note{
		Title:    "",
		Comments: "checkout Nasa.gov",
		DueDate:  "2019-10-21",
	}
	res, err := s.server.CreateNote(ctx, &pb.CreateNoteMsg{
		Note: note,
	})
	s.Assert().Nil(res)
	s.Assert().Error(err)
	s.Assert().EqualError(err, utils.ErrEmptyTitle.Error())
}

func (s *Suite) Test03CreateNoteInvalidTimeFormat() {
	note := &pb.Note{
		Title:    "Meteor Shower",
		Comments: "checkout Nasa.gov",
		DueDate:  "0-0-0",
	}
	res, err := s.server.CreateNote(ctx, &pb.CreateNoteMsg{
		Note: note,
	})
	s.Assert().Nil(res)
	s.Assert().Error(err)
	s.Assert().EqualError(err, utils.ErrInvalidTime.Error())
}

func (s *Suite) Test04GetAllNotes() {
	res, err := s.server.GetAllNotes(ctx, &pb.Empty{})
	s.Assert().NoError(err)
	notes := res.Notes
	s.Assert().Equal(1, len(notes))

	expectedNote := &pb.Note{
		Title:    "Meteor Shower",
		Comments: "checkout Nasa.gov",
		DueDate:  "2019-10-21",
	}
	gotNote := notes[0]
	s.Assert().Equal(expectedNote.Title, gotNote.Title)
	s.Assert().Equal(expectedNote.Comments, gotNote.Comments)
	s.Assert().Equal(expectedNote.DueDate, gotNote.DueDate)
	s.Assert().NotZero(gotNote.Id)
}

func (s *Suite) Test06UpdateNote() {
	allNotes, err := s.server.GetAllNotes(ctx, &pb.Empty{})
	s.Assert().NoError(err)
	noteId := allNotes.Notes[0].Id
	updatedNote := pb.Note{
		Id:       noteId,
		Title:    "New update Meteor Shower",
		Comments: "checkout Nasa.gov",
		DueDate:  "2019-10-21",
	}
	res, err := s.server.UpdateNote(ctx, &pb.UpdateNoteMsg{
		Note: &updatedNote,
	})
	s.Assert().NoError(err)
	s.Assert().NotNil(res)
	s.Assert().Equal(res.Note.Id, updatedNote.Id)
	s.Assert().Equal(res.Note.Comments, updatedNote.Comments)
	s.Assert().Equal(res.Note.DueDate, updatedNote.DueDate)

}

func (s *Suite) Test07UpdateInvalidNote() {
	updatedNote := pb.Note{
		Id:       0,
		Title:    "New update Meteor Shower",
		Comments: "checkout Nasa.gov",
		DueDate:  "2019-10-21",
	}
	res, err := s.server.UpdateNote(ctx, &pb.UpdateNoteMsg{
		Note: &updatedNote,
	})

	s.Assert().Nil(res)
	s.Assert().EqualError(err, utils.ErrEmptyNote.Error())
}

func (s *Suite) Test08DeleteNote() {
	res, err := s.server.GetAllNotes(ctx, &pb.Empty{})
	s.Assert().NoError(err)
	notes := res.Notes
	s.Assert().Equal(1, len(notes))
	gotNote := notes[0]

	deleteRes, err := s.server.DeleteNote(ctx, &pb.DeleteNoteMsg{
		Id: gotNote.Id,
	})
	s.Assert().NoError(err)
	s.Assert().Equal(&pb.Empty{}, deleteRes)
}

func (s *Suite) Test09DeleteNoteWithInvalidNoteId() {
	res, err := s.server.DeleteNote(ctx, &pb.DeleteNoteMsg{
		Id: 0,
	})
	s.Assert().Nil(res)
	s.Assert().EqualError(err, utils.ErrEmptyNote.Error())

	res, err = s.server.DeleteNote(ctx, &pb.DeleteNoteMsg{
		Id: 100,
	})
	s.Assert().Nil(res)
	s.Assert().EqualError(err,
		errors.WithMessagef(utils.ErrEmptyNotes, "could not delete note with id: %d", 100).Error())
}
func (s *Suite) Test10DeleteNoteWhenNotesListIsEmpty() {
	res, err := s.server.GetAllNotes(ctx, &pb.Empty{})
	s.Assert().NoError(err)
	notes := res.Notes
	s.Assert().Equal(0, len(notes))

	deleteRes, err := s.server.DeleteNote(ctx, &pb.DeleteNoteMsg{
		Id: 1,
	})
	s.Assert().Nil(deleteRes)
	s.Assert().EqualError(err,
		errors.WithMessagef(utils.ErrEmptyNotes, "could not delete note with id: %d", 1).Error())
}

func TestSuite(t *testing.T) {
	testSuite := new(Suite)
	testSuite.server = NewTodoServer(repo.NewTodoService())
	suite.Run(t, testSuite)
}
