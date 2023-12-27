package note

import (
	"context"

	"github.com/ansedo/note-service-api/internal/model"
)

type Service struct {
	noteRepository noteRepository
}

type noteRepository interface {
	Create(ctx context.Context, note *model.Note) (*model.Note, error)
	Note(ctx context.Context, req *model.Note) (*model.Note, error)
	Notes(ctx context.Context) ([]*model.Note, error)
	Update(ctx context.Context, note *model.Note) error
	Delete(ctx context.Context, note *model.Note) error
}

func NewService(noteRepository noteRepository) *Service {
	return &Service{
		noteRepository: noteRepository,
	}
}
