package repository

import (
	"context"

	sq "github.com/Masterminds/squirrel"
	"github.com/jmoiron/sqlx"

	"github.com/ansedo/note-service-api/internal/model"
	"github.com/ansedo/note-service-api/internal/repository/table"
)

type NoteRepository interface {
	Create(ctx context.Context, note *model.Note) (*model.Note, error)
	Note(ctx context.Context, req *model.Note) (*model.Note, error)
	Notes(ctx context.Context) ([]*model.Note, error)
	Update(ctx context.Context, note *model.Note) error
	Delete(ctx context.Context, note *model.Note) error
}

type Repository struct {
	db *sqlx.DB
}

func NewNoteRepository(db *sqlx.DB) NoteRepository {
	return &Repository{
		db: db,
	}
}

func (r *Repository) Create(ctx context.Context, note *model.Note) (*model.Note, error) {
	query, args, err := sq.Insert(table.Note).
		PlaceholderFormat(sq.Dollar).
		Columns(table.ColumnTitle, table.ColumnText, table.ColumnAuthor, table.ColumnEmail).
		Values(note.Title, note.Text, note.Author, note.Email).
		Suffix(table.InsertSuffix).
		ToSql()
	if err != nil {
		return nil, err
	}

	row, err := r.db.QueryContext(ctx, query, args...)
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

	return &model.Note{Id: id}, nil
}
func (r *Repository) Note(ctx context.Context, note *model.Note) (*model.Note, error) {
	query, args, err := sq.Select(table.ColumnTitle, table.ColumnText, table.ColumnAuthor, table.ColumnEmail).
		PlaceholderFormat(sq.Dollar).
		From(table.Note).
		Where(sq.Eq{table.ColumnId: note.Id}).
		ToSql()
	if err != nil {
		return nil, err
	}

	row, err := r.db.QueryContext(ctx, query, args...)
	if err != nil {
		return nil, err
	}
	defer row.Close()

	var title, text, author, email string
	row.Next()
	if err = row.Scan(&title, &text, &author, &email); err != nil {
		return nil, err
	}

	return &model.Note{
		Id:     note.Id,
		Title:  title,
		Text:   text,
		Author: author,
		Email:  email,
	}, nil
}

func (r *Repository) Notes(ctx context.Context) ([]*model.Note, error) {
	query, args, err := sq.Select(table.ColumnId, table.ColumnTitle, table.ColumnText, table.ColumnAuthor, table.ColumnEmail).
		PlaceholderFormat(sq.Dollar).
		From(table.Note).
		ToSql()
	if err != nil {
		return nil, err
	}

	row, err := r.db.QueryContext(ctx, query, args...)
	if err != nil {
		return nil, err
	}
	defer row.Close()

	var (
		id                         int64
		title, text, author, email string
	)
	var notes []*model.Note
	for row.Next() {
		if err = row.Scan(&id, &title, &text, &author, &email); err != nil {
			return nil, err
		}
		notes = append(
			notes,
			&model.Note{
				Id:     id,
				Title:  title,
				Text:   text,
				Author: author,
				Email:  email,
			},
		)
	}

	return notes, nil
}

func (r *Repository) Update(ctx context.Context, note *model.Note) error {
	clauses := make(map[string]interface{})
	clauses[table.ColumnUpdatedAt] = sq.Expr("NOW()")

	if note.Title != "" {
		clauses[table.ColumnTitle] = note.Title
	}
	if note.Text != "" {
		clauses[table.ColumnText] = note.Text
	}
	if note.Author != "" {
		clauses[table.ColumnAuthor] = note.Author
	}
	if note.Email != "" {
		clauses[table.ColumnEmail] = note.Email
	}

	query, args, err := sq.Update(table.Note).
		PlaceholderFormat(sq.Dollar).
		SetMap(clauses).
		Where(sq.Eq{table.ColumnId: note.Id}).
		ToSql()
	if err != nil {
		return err
	}

	if _, err = r.db.ExecContext(ctx, query, args...); err != nil {
		return err
	}

	return nil
}

func (r *Repository) Delete(ctx context.Context, note *model.Note) error {
	query, args, err := sq.Delete(table.Note).
		PlaceholderFormat(sq.Dollar).
		Where(sq.Eq{table.ColumnId: note.Id}).
		ToSql()
	if err != nil {
		return err
	}

	if _, err = r.db.ExecContext(ctx, query, args...); err != nil {
		return err
	}

	return nil
}
