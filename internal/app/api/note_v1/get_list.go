package note_v1

import (
	"context"
	"log/slog"

	"github.com/golang/protobuf/ptypes/empty"

	desc "github.com/ansedo/note-service-api/pkg/note_v1"
)

func (n *Note) GetList(ctx context.Context, req *empty.Empty) (*desc.GetListResponse, error) {
	slog.Info(
		"method `GetList` has been called",
		slog.String("op", "app.api.note_v1.Notes"),
	)

	notes, err := n.noteService.Notes(ctx)
	if err != nil {
		return nil, err
	}

	descNotes := make([]*desc.Note, 0, len(notes))
	for _, note := range notes {
		descNotes = append(descNotes, note.ToDescNote())
	}

	return &desc.GetListResponse{Notes: descNotes}, nil
}
