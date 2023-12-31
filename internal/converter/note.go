package converter

import (
	"database/sql"

	"google.golang.org/protobuf/types/known/timestamppb"

	"github.com/ansedo/note-service-api/internal/model"
	desc "github.com/ansedo/note-service-api/pkg/note_v1"
)

func ToNoteInfo(noteInfo *desc.NoteInfo) *model.NoteInfo {
	return &model.NoteInfo{
		Title:  noteInfo.GetTitle(),
		Text:   noteInfo.GetText(),
		Author: noteInfo.GetAuthor(),
		Email:  noteInfo.GetEmail(),
	}
}

func ToUpdateNoteInfo(updateNoteInfo *desc.UpdateNoteInfo) *model.UpdateNoteInfo {
	res := &model.UpdateNoteInfo{}
	if updateNoteInfo.Title != nil {
		res.Title = sql.NullString{String: updateNoteInfo.GetTitle().String(), Valid: true}
	}
	if updateNoteInfo.Text != nil {
		res.Text = sql.NullString{String: updateNoteInfo.GetText().String(), Valid: true}
	}
	if updateNoteInfo.Author != nil {
		res.Author = sql.NullString{String: updateNoteInfo.GetAuthor().String(), Valid: true}
	}
	if updateNoteInfo.Email != nil {
		res.Email = sql.NullString{String: updateNoteInfo.GetEmail().String(), Valid: true}
	}

	return res
}

func ToDescNoteInfo(noteInfo *model.NoteInfo) *desc.NoteInfo {
	return &desc.NoteInfo{
		Title:  noteInfo.Title,
		Text:   noteInfo.Text,
		Author: noteInfo.Author,
		Email:  noteInfo.Email,
	}
}

func ToDescNote(note *model.Note) *desc.Note {
	var updatedAt *timestamppb.Timestamp
	if note.UpdatedAt.Valid {
		updatedAt = timestamppb.New(note.UpdatedAt.Time)
	}

	return &desc.Note{
		Id:        note.ID,
		Info:      ToDescNoteInfo(note.Info),
		CreatedAt: timestamppb.New(note.CreatedAt),
		UpdatedAt: updatedAt,
	}
}
