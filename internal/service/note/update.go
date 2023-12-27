package note

import (
	"context"

	"github.com/ansedo/note-service-api/internal/model"
)

func (s *Service) Update(ctx context.Context, note *model.Note) error {
	return s.noteRepository.Update(ctx, note)
}
