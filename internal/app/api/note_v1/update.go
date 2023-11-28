package note_v1

import (
	"context"
	"log/slog"

	desc "github.com/ansedo/note-service-api/pkg/note_v1"
	"github.com/golang/protobuf/ptypes/empty"
)

func (n *Note) Update(ctx context.Context, req *desc.UpdateRequest) (*empty.Empty, error) {
	slog.Info(
		"method `Update` has been called",
		slog.String("op", "app.api.note_v1.Update"),
		slog.Any("request", req),
	)

	return &empty.Empty{}, nil
}
