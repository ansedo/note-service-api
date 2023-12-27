package note

import (
	"context"

	"github.com/ansedo/note-service-api/internal/model"
)

func (s *Service) GetList(ctx context.Context) ([]*model.Note, error) {
	return s.noteRepository.GetList(ctx)
}
