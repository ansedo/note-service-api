package model

import (
	"database/sql"
	"time"
)

type NoteInfo struct {
	Title  string `db:"title"`
	Text   string `db:"text"`
	Author string `db:"author"`
	Email  string `db:"email"`
}

type Note struct {
	ID        int64        `db:"id"`
	Info      *NoteInfo    `db:""`
	CreatedAt time.Time    `db:"created_at"`
	UpdatedAt sql.NullTime `db:"updated_at"`
}

type UpdateNoteInfo struct {
	Title  sql.NullString `db:"title"`
	Text   sql.NullString `db:"text"`
	Author sql.NullString `db:"author"`
	Email  sql.NullString `db:"email"`
}
