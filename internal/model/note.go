package model

import desc "github.com/ansedo/note-service-api/pkg/note_v1"

type Note struct {
	Id     int64  `db:"id"`
	Title  string `db:"title"`
	Text   string `db:"text"`
	Author string `db:"author"`
	Email  string `db:"email"`
}

func NewNoteFromDesc(note *desc.Note) *Note {
	return &Note{
		Id:     note.GetId(),
		Title:  note.GetTitle(),
		Text:   note.GetText(),
		Author: note.GetAuthor(),
		Email:  note.GetEmail(),
	}
}

func (n *Note) ToDescNote() *desc.Note {
	return &desc.Note{
		Id:     n.Id,
		Title:  n.Title,
		Text:   n.Text,
		Author: n.Author,
		Email:  n.Email,
	}
}
