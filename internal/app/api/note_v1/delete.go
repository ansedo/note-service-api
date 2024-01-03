package note_v1

import (
	"context"

	"github.com/golang/protobuf/ptypes/empty"

	desc "github.com/ansedo/note-service-api/pkg/note_v1"
)

func (n *Note) Delete(ctx context.Context, req *desc.DeleteRequest) (*empty.Empty, error) {
	if err := n.noteService.Delete(ctx, req.GetId()); err != nil {
		return nil, err
	}

	return &empty.Empty{}, nil
}
