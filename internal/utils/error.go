package utils

import (
	"errors"
)

const DateLayout = "2006-01-02"

var (
	ErrEmptyTitle  = errors.New("title is required")
	ErrNotFound    = errors.New("note does not exist")
	ErrEmptyNotes  = errors.New("no notes found")
	ErrEmptyNote   = errors.New("valid note required but received empty")
	ErrInvalidTime = errors.New("invalid time")
)
