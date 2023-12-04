package note_v1

import (
	"context"
	"log/slog"

	sq "github.com/Masterminds/squirrel"
	_ "github.com/jackc/pgx/stdlib"
	"github.com/jmoiron/sqlx"

	desc "github.com/ansedo/note-service-api/pkg/note_v1"
)

const (
	sqlInsertSuffix = "RETURNING ID"
)

func (n *Note) Create(ctx context.Context, req *desc.CreateRequest) (*desc.CreateResponse, error) {
	slog.Info(
		"method `Create` has been called",
		slog.String("op", "app.api.note_v1.create"),
		slog.Any("request", req),
	)

	db, err := sqlx.Open("pgx", dbDsn)
	if err != nil {
		return nil, err
	}
	defer db.Close()

	query, args, err := sq.Insert(noteTable).
		PlaceholderFormat(sq.Dollar).
		Columns(sqlColumnTitle, sqlColumnText, sqlColumnAuthor).
		Values(req.GetTitle(), req.GetText(), req.GetAuthor()).
		Suffix(sqlInsertSuffix).
		ToSql()
	if err != nil {
		return nil, err
	}

	row, err := db.QueryContext(ctx, query, args...)
	if err != nil {
		return nil, err
	}
	defer row.Close()

	var id int64
	row.Next()
	err = row.Scan(&id)
	if err != nil {
		return nil, err
	}

	return &desc.CreateResponse{
		Id: id,
	}, nil
}
