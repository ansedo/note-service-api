package note_v1

import (
	"context"

	"github.com/ansedo/note-service-api/internal/converter"
	"github.com/golang/protobuf/ptypes/empty"

	desc "github.com/ansedo/note-service-api/pkg/note_v1"
)

func (n *Note) GetList(ctx context.Context, req *empty.Empty) (*desc.GetListResponse, error) {
	notes, err := n.noteService.GetList(ctx)
	if err != nil {
		return nil, err
	}

	descNotes := make([]*desc.Note, 0, len(notes))
	for _, note := range notes {
		descNotes = append(descNotes, converter.ToDescNote(note))
	}

	return &desc.GetListResponse{Notes: descNotes}, nil
}
