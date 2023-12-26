package note_v1

import (
	"context"

	"github.com/ansedo/note-service-api/internal/model"
	desc "github.com/ansedo/note-service-api/pkg/note_v1"
)

func (n *Note) Get(ctx context.Context, req *desc.GetRequest) (*desc.GetResponse, error) {
	note, err := n.noteService.Note(ctx, &model.Note{Id: req.GetId()})
	if err != nil {
		return nil, err
	}

	return &desc.GetResponse{
		Note: note.ToDescNote(),
	}, nil
}
