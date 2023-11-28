package note_v1

import (
	"context"
	"log/slog"

	desc "github.com/ansedo/note-service-api/pkg/note_v1"
	"github.com/golang/protobuf/ptypes/empty"
)

func (n *Note) Delete(ctx context.Context, req *desc.DeleteRequest) (*empty.Empty, error) {
	slog.Info(
		"method `Delete` has been called",
		slog.String("op", "app.api.note_v1.Delete"),
		slog.Any("request", req),
	)

	return &empty.Empty{}, nil
}
