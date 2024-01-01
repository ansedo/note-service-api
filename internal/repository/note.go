package repository

import (
	"context"

	sq "github.com/Masterminds/squirrel"
	"github.com/ansedo/note-service-api/internal/model"
	"github.com/ansedo/note-service-api/internal/pkg/db"
	"github.com/ansedo/note-service-api/internal/repository/table"
)

type NoteRepository struct {
	client *db.Client
}

func NewNoteRepository(client *db.Client) *NoteRepository {
	return &NoteRepository{
		client: client,
	}
}

func (r *NoteRepository) Create(ctx context.Context, noteInfo *model.NoteInfo) (int64, error) {
	query, args, err := sq.Insert(table.Note).
		PlaceholderFormat(sq.Dollar).
		Columns(table.ColumnTitle, table.ColumnText, table.ColumnAuthor, table.ColumnEmail).
		Values(noteInfo.Title, noteInfo.Text, noteInfo.Author, noteInfo.Email).
		Suffix(table.InsertSuffix).
		ToSql()
	if err != nil {
		return 0, err
	}

	q := db.Query{
		Name:     "CreateNote",
		QueryRaw: query,
	}

	row, err := r.client.DB().Query(ctx, q, args...)
	if err != nil {
		return 0, err
	}
	defer row.Close()

	var id int64
	row.Next()
	err = row.Scan(&id)
	if err != nil {
		return 0, err
	}

	return id, nil
}

func (r *NoteRepository) Get(ctx context.Context, id int64) (*model.Note, error) {
	query, args, err := sq.Select(table.ColumnTitle, table.ColumnText, table.ColumnAuthor, table.ColumnEmail).
		PlaceholderFormat(sq.Dollar).
		From(table.Note).
		Where(sq.Eq{table.ColumnId: id}).
		ToSql()
	if err != nil {
		return nil, err
	}

	q := db.Query{
		Name:     "GetNote",
		QueryRaw: query,
	}

	var note model.Note
	if err = r.client.DB().Get(ctx, &note, q, args...); err != nil {
		return nil, err
	}

	return &note, nil
}

func (r *NoteRepository) GetList(ctx context.Context) ([]*model.Note, error) {
	query, args, err := sq.Select(table.ColumnId, table.ColumnTitle, table.ColumnText, table.ColumnAuthor, table.ColumnEmail).
		PlaceholderFormat(sq.Dollar).
		From(table.Note).
		ToSql()
	if err != nil {
		return nil, err
	}

	q := db.Query{
		Name:     "GetList",
		QueryRaw: query,
	}

	var notes []*model.Note
	if err = r.client.DB().Select(ctx, &notes, q, args...); err != nil {
		return nil, err
	}

	return notes, nil
}

func (r *NoteRepository) Update(ctx context.Context, id int64, updateNoteInfo *model.UpdateNoteInfo) error {
	builder := sq.Update(table.Note).
		Set(table.ColumnUpdatedAt, sq.Expr("NOW()")).
		PlaceholderFormat(sq.Dollar).
		Where(sq.Eq{table.ColumnId: id})

	if updateNoteInfo.Title.Valid {
		builder = builder.Set(table.ColumnTitle, updateNoteInfo.Title.String)
	}
	if updateNoteInfo.Text.Valid {
		builder = builder.Set(table.ColumnText, updateNoteInfo.Text.String)
	}
	if updateNoteInfo.Author.Valid {
		builder = builder.Set(table.ColumnAuthor, updateNoteInfo.Author.String)
	}
	if updateNoteInfo.Email.Valid {
		builder = builder.Set(table.ColumnEmail, updateNoteInfo.Email.String)
	}

	query, args, err := builder.ToSql()
	if err != nil {
		return err
	}

	q := db.Query{
		Name:     "UpdateNote",
		QueryRaw: query,
	}

	if _, err = r.client.DB().Exec(ctx, q, args...); err != nil {
		return err
	}

	return nil
}

func (r *NoteRepository) Delete(ctx context.Context, id int64) error {
	query, args, err := sq.Delete(table.Note).
		PlaceholderFormat(sq.Dollar).
		Where(sq.Eq{table.ColumnId: id}).
		ToSql()
	if err != nil {
		return err
	}

	q := db.Query{
		Name:     "DeleteNote",
		QueryRaw: query,
	}

	if _, err = r.client.DB().Exec(ctx, q, args...); err != nil {
		return err
	}

	return nil
}
