package note

import (
	"context"

	"github.com/ansedo/note-service-api/internal/model"
)

func (s *Service) Notes(ctx context.Context) ([]*model.Note, error) {
	return s.noteRepository.Notes(ctx)
}
