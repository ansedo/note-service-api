package note

import (
	"context"

	"github.com/ansedo/note-service-api/internal/model"
)

type Service struct {
	noteRepository noteRepository
}

type noteRepository interface {
	Create(ctx context.Context, noteInfo *model.NoteInfo) (int64, error)
	Get(ctx context.Context, id int64) (*model.Note, error)
	GetList(ctx context.Context) ([]*model.Note, error)
	Update(ctx context.Context, id int64, updateNoteInfo *model.UpdateNoteInfo) error
	Delete(ctx context.Context, id int64) error
}

func NewService(noteRepository noteRepository) *Service {
	return &Service{
		noteRepository: noteRepository,
	}
}

func NewMockNoteService(deps ...any) *Service {
	s := &Service{}

	for _, v := range deps {
		switch t := v.(type) {
		case noteRepository:
			s.noteRepository = t
		}
	}

	return s
}
