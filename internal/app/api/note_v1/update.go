package note_v1

import (
	"context"
	"log/slog"

	"github.com/ansedo/note-service-api/internal/model"
	"github.com/golang/protobuf/ptypes/empty"

	desc "github.com/ansedo/note-service-api/pkg/note_v1"
)

func (n *Note) Update(ctx context.Context, req *desc.UpdateRequest) (*empty.Empty, error) {
	slog.Info(
		"method `Update` has been called",
		slog.String("op", "app.api.note_v1.Update"),
		slog.Any("request", req),
	)

	return &empty.Empty{}, n.noteService.Update(ctx, model.NewNoteFromDesc(req.GetNote()))
}
