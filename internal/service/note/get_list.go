package note

import (
	"context"

	"github.com/ansedo/note-service-api/internal/model"
)

func (s *Service) GetList(ctx context.Context) ([]*model.Note, error) {
	notes, err := s.noteRepository.GetList(ctx)
	if err != nil {
		return nil, err
	}

	return notes, nil
}
