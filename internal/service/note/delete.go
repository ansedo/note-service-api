package note

import (
	"context"
)

func (s *Service) Delete(ctx context.Context, id int64) error {
	return s.noteRepository.Delete(ctx, id)
}
