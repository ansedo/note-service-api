package note_v1

import (
	"context"
	"log/slog"

	sq "github.com/Masterminds/squirrel"
	_ "github.com/jackc/pgx/stdlib"
	"github.com/jmoiron/sqlx"

	desc "github.com/ansedo/note-service-api/pkg/note_v1"
)

func (n *Note) Get(ctx context.Context, req *desc.GetRequest) (*desc.GetResponse, error) {
	slog.Info(
		"method `Get` has been called",
		slog.String("op", "app.api.note_v1.Get"),
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
		Where(sq.Eq{sqlColumnId: req.GetId()}).
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
	row.Next()
	if err = row.Scan(&id, &title, &text, &author); err != nil {
		return nil, err
	}

	return &desc.GetResponse{
		Note: &desc.Note{
			Id:     id,
			Title:  title,
			Text:   text,
			Author: author,
		},
	}, nil
}
