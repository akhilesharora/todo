package server

import (
	"context"
	"math/rand"
	"time"

	"github.com/pkg/errors"

	"github.com/akhilesharora/todo/internal/repo"
	"github.com/akhilesharora/todo/internal/utils"
	"github.com/akhilesharora/todo/pb"
)

type Server struct {
	repo repo.Dashboard
}

func NewTodoServer(repo repo.Dashboard) *Server {
	return &Server{
		repo: repo,
	}
}

func (s *Server) CreateNote(ctx context.Context, in *pb.CreateNoteMsg) (*pb.Empty, error) {
	if in.Note == nil {
		return nil, utils.ErrEmptyNote
	}

	if in.Note.Title == "" {
		return nil, utils.ErrEmptyTitle
	}

	t, err := time.Parse(utils.DateLayout, in.Note.DueDate)
	if err != nil {
		return nil, utils.ErrInvalidTime
	}

	err = s.repo.Create(&repo.Note{
		ID:       rand.Uint32(),
		Title:    in.Note.Title,
		Comments: in.Note.Comments,
		DueDate:  t,
	})
	if err != nil {
		return &pb.Empty{}, err
	}

	return &pb.Empty{}, nil
}

func (s *Server) GetAllNotes(ctx context.Context, in *pb.Empty) (*pb.GetAllNotesReply, error) {
	notes, err := s.repo.GetAll()
	if err != nil {
		return &pb.GetAllNotesReply{}, nil
	}

	allNotes := make([]*pb.Note, 0)
	for _, note := range notes {
		allNotes = append(allNotes, &pb.Note{
			Id:       note.ID,
			Title:    note.Title,
			Comments: note.Comments,
			DueDate:  note.DueDate.Format(utils.DateLayout),
		})
	}

	return &pb.GetAllNotesReply{
		Notes: allNotes,
	}, nil
}

func (s *Server) UpdateNote(ctx context.Context, in *pb.UpdateNoteMsg) (*pb.UpdateNoteReply, error) {
	if in.Note == nil || in.Note.Id == 0 {
		return nil, utils.ErrEmptyNote
	}

	if in.Note.Title == "" {
		return nil, utils.ErrEmptyTitle
	}

	t, err := time.Parse(utils.DateLayout, in.Note.DueDate)
	if err != nil {
		return nil, utils.ErrInvalidTime
	}

	note, err := s.repo.Update(&repo.Note{
		ID:       in.Note.Id,
		Title:    in.Note.Title,
		Comments: in.Note.Comments,
		DueDate:  t,
	})

	if err != nil {
		return nil, err
	}

	return &pb.UpdateNoteReply{
		Note: &pb.Note{
			Id:       note.ID,
			Title:    note.Title,
			Comments: note.Comments,
			DueDate:  note.DueDate.Format(utils.DateLayout),
		},
	}, nil
}

func (s *Server) DeleteNote(ctx context.Context, in *pb.DeleteNoteMsg) (*pb.Empty, error) {
	if in.Id == 0 {
		return nil, utils.ErrEmptyNote
	}
	err := s.repo.Delete(in.Id)
	if err != nil {
		return nil, errors.WithMessagef(err, "could not delete note with id: %d", in.Id)
	}
	return &pb.Empty{}, nil
}
