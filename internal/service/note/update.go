package note

import (
	"context"

	"github.com/ansedo/note-service-api/internal/model"
)

func (s *Service) Update(ctx context.Context, id int64, updateNoteInfo *model.UpdateNoteInfo) error {
	return s.noteRepository.Update(ctx, id, updateNoteInfo)
}
