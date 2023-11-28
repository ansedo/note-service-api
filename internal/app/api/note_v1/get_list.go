package note_v1

import (
	"context"
	"log/slog"

	desc "github.com/ansedo/note-service-api/pkg/note_v1"
	"github.com/golang/protobuf/ptypes/empty"
)

func (n *Note) GetList(ctx context.Context, req *empty.Empty) (*desc.GetListResponse, error) {
	slog.Info(
		"method `GetList` has been called",
		slog.String("op", "app.api.note_v1.GetList"),
		slog.Any("request", req),
	)

	return &desc.GetListResponse{
		Notes: []*desc.Note{
			{
				Id:     1,
				Title:  "Title #1 you got",
				Text:   "Text #1 you got",
				Author: "Author #1 you got",
			},
			{
				Id:     2,
				Title:  "Title #2 you got",
				Text:   "Text #2 you got",
				Author: "Author #2 you got",
			},
			{
				Id:     3,
				Title:  "Title #3 you got",
				Text:   "Text #3 you got",
				Author: "Author #3 you got",
			},
		},
	}, nil
}
