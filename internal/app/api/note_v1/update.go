package note_v1

import (
	"context"

	"github.com/ansedo/note-service-api/internal/converter"
	"github.com/golang/protobuf/ptypes/empty"

	desc "github.com/ansedo/note-service-api/pkg/note_v1"
)

func (n *Note) Update(ctx context.Context, req *desc.UpdateRequest) (*empty.Empty, error) {
	if err := n.noteService.Update(ctx, req.GetId(), converter.ToUpdateNoteInfo(req.GetNote())); err != nil {
		return nil, err
	}

	return &empty.Empty{}, nil
}
