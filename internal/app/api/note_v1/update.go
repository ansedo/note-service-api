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

func (n *Note) Update(ctx context.Context, req *desc.UpdateRequest) (*empty.Empty, error) {
	slog.Info(
		"method `Update` has been called",
		slog.String("op", "app.api.note_v1.Update"),
		slog.Any("request", req),
	)

	db, err := sqlx.Open("pgx", dbDsn)
	if err != nil {
		return nil, err
	}
	defer db.Close()

	query, args, err := sq.Update(noteTable).
		PlaceholderFormat(sq.Dollar).
		Set(sqlColumnTitle, req.Note.GetTitle()).
		Set(sqlColumnText, req.Note.GetText()).
		Set(sqlColumnAuthor, req.Note.GetAuthor()).
		Set(sqlColumnUpdatedAt, sq.Expr("NOW()")).
		Where(sq.Eq{sqlColumnId: req.Note.Id}).
		ToSql()
	if err != nil {
		return nil, err
	}

	if _, err = db.ExecContext(ctx, query, args...); err != nil {
		return nil, err
	}

	return &empty.Empty{}, nil
}
