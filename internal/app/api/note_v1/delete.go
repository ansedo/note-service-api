package note_v1

import (
	"context"

	"github.com/golang/protobuf/ptypes/empty"

	"github.com/ansedo/note-service-api/internal/model"
	desc "github.com/ansedo/note-service-api/pkg/note_v1"
)

func (n *Note) Delete(ctx context.Context, req *desc.DeleteRequest) (*empty.Empty, error) {
	return &empty.Empty{}, n.noteService.Delete(ctx, &model.Note{Id: req.GetId()})
}
