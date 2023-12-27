package note

import (
	"context"

	"github.com/ansedo/note-service-api/internal/model"
)

func (s *Service) Get(ctx context.Context, note *model.Note) (*model.Note, error) {
	return s.noteRepository.Get(ctx, note)
}
