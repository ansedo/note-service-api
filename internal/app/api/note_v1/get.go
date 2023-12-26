package note_v1

import (
	"context"
	"log/slog"

	"github.com/ansedo/note-service-api/internal/model"
	desc "github.com/ansedo/note-service-api/pkg/note_v1"
)

func (n *Note) Get(ctx context.Context, req *desc.GetRequest) (*desc.GetResponse, error) {
	slog.Info(
		"method `Get` has been called",
		slog.String("op", "app.api.note_v1.Get"),
		slog.Any("req", req),
	)

	note, err := n.noteService.Note(ctx, &model.Note{Id: req.GetId()})
	if err != nil {
		return nil, err
	}

	return &desc.GetResponse{
		Note: note.ToDescNote(),
	}, nil
}
