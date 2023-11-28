package note_v1

import (
	"context"
	"log/slog"

	desc "github.com/ansedo/note-service-api/pkg/note_v1"
)

func (n *Note) Get(ctx context.Context, req *desc.GetRequest) (*desc.GetResponse, error) {
	slog.Info(
		"method `Get` has been called",
		slog.String("op", "app.api.note_v1.Get"),
		slog.Any("request", req),
	)

	return &desc.GetResponse{
		Note: &desc.Note{
			Id:     1,
			Title:  "Title you got",
			Text:   "Text you got",
			Author: "Author you got",
		},
	}, nil
}
