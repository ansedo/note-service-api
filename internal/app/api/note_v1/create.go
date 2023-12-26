package note_v1

import (
	"context"

	"github.com/ansedo/note-service-api/internal/model"
	desc "github.com/ansedo/note-service-api/pkg/note_v1"
)

func (n *Note) Create(ctx context.Context, req *desc.CreateRequest) (*desc.CreateResponse, error) {
	note, err := n.noteService.Create(
		ctx,
		&model.Note{
			Title:  req.GetTitle(),
			Text:   req.GetText(),
			Author: req.GetAuthor(),
			Email:  req.GetEmail(),
		},
	)
	if err != nil {
		return nil, err
	}

	return &desc.CreateResponse{Id: note.Id}, nil
}
