package note_v1

import (
	"context"

	"github.com/golang/protobuf/ptypes/empty"

	"github.com/ansedo/note-service-api/internal/model"
	desc "github.com/ansedo/note-service-api/pkg/note_v1"
)

func (n *Note) Update(ctx context.Context, req *desc.UpdateRequest) (*empty.Empty, error) {
	return &empty.Empty{}, n.noteService.Update(ctx, model.NewNoteFromDesc(req.GetNote()))
}
