package repo

import (
	"time"

	"github.com/akhilesharora/todo/internal/utils"
)

type Note struct {
	ID       uint32    `json:"id"`
	Title    string    `json:"title"`
	Comments string    `json:"comments"`
	DueDate  time.Time `json:"due_date"`
}

type Dashboard interface {
	Create(note *Note) error
	GetAll() ([]*Note, error)
	Update(note *Note) (*Note, error)
	Delete(noteId uint32) error
}

type TodoService struct {
	notes []*Note
}

func NewTodoService() *TodoService {
	return &TodoService{
		notes: []*Note{},
	}
}

func (s *TodoService) Create(note *Note) error {
	if note == nil {
		return utils.ErrEmptyNote
	}
	s.notes = append(s.notes, note)
	return nil
}

func (s *TodoService) get(noteId uint32) (*Note, error) {
	if len(s.notes) == 0 {
		return nil, utils.ErrEmptyNotes
	}
	for _, v := range s.notes {
		if v.ID == noteId {
			return v, nil
		}
	}
	return nil, utils.ErrNotFound
}

func (s *TodoService) GetAll() ([]*Note, error) {
	if len(s.notes) == 0 {
		return []*Note{}, utils.ErrEmptyNotes
	}
	return s.notes, nil
}

func (s *TodoService) Update(note *Note) (*Note, error) {
	oldNote, err := s.get(note.ID)
	if err != nil {
		return nil, err
	}
	err = s.Delete(oldNote.ID)
	if err != nil {
		return nil, err
	}
	s.notes = append(s.notes, note)
	return note, nil
}

func (s *TodoService) Delete(noteId uint32) error {
	length := len(s.notes)
	if length == 0 {
		return utils.ErrEmptyNotes
	}
	flag := false
	for k, v := range s.notes {
		if v.ID == noteId {
			flag = true
			s.notes[k] = s.notes[length-1]
			s.notes[length-1] = &Note{}
			s.notes = s.notes[:length-1]
		}
	}
	if !flag {
		return utils.ErrNotFound
	}
	return nil
}
