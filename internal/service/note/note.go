package note

import (
	"context"

	"github.com/ansedo/note-service-api/internal/model"
)

func (s *Service) Note(ctx context.Context, note *model.Note) (*model.Note, error) {
	return s.noteRepository.Note(ctx, note)
}
