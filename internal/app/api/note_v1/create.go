package note_v1

import (
	"context"
	"log/slog"

	desc "github.com/ansedo/note-service-api/pkg/note_v1"
)

func (n *Note) Create(ctx context.Context, req *desc.CreateRequest) (*desc.CreateResponse, error) {
	slog.Info(
		"method `Create` has been called",
		slog.String("op", "app.api.note_v1.create"),
		slog.Any("request", req),
	)

	return &desc.CreateResponse{
		Id: 1,
	}, nil
}
