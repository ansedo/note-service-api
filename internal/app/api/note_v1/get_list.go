package note_v1

import (
	"context"
	"log/slog"

	sq "github.com/Masterminds/squirrel"
	"github.com/golang/protobuf/ptypes/empty"
	_ "github.com/jackc/pgx/stdlib"
	"github.com/jmoiron/sqlx"

	desc "github.com/ansedo/note-service-api/pkg/note_v1"
)

func (n *Note) GetList(ctx context.Context, req *empty.Empty) (*desc.GetListResponse, error) {
	slog.Info(
		"method `GetList` has been called",
		slog.String("op", "app.api.note_v1.GetList"),
		slog.Any("request", req),
	)

	db, err := sqlx.Open("pgx", dbDsn)
	if err != nil {
		return nil, err
	}
	defer db.Close()

	query, args, err := sq.Select(sqlColumnId, sqlColumnTitle, sqlColumnText, sqlColumnAuthor).
		PlaceholderFormat(sq.Dollar).
		From(noteTable).
		ToSql()
	if err != nil {
		return nil, err
	}

	row, err := db.QueryContext(ctx, query, args...)
	if err != nil {
		return nil, err
	}
	defer row.Close()

	var (
		id                  int64
		title, text, author string
	)
	var notes []*desc.Note
	for row.Next() {
		if err = row.Scan(&id, &title, &text, &author); err != nil {
			return nil, err
		}
		notes = append(
			notes,
			&desc.Note{
				Id:     id,
				Title:  title,
				Text:   text,
				Author: author,
			},
		)
	}

	return &desc.GetListResponse{Notes: notes}, nil
}
