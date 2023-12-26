package note

import (
	"context"

	"github.com/ansedo/note-service-api/internal/model"
)

func (s *Service) Delete(ctx context.Context, note *model.Note) error {
	return s.noteRepository.Delete(ctx, note)
}
