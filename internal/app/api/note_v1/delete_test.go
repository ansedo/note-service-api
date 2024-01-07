package note_v1

import (
	"context"
	"errors"
	"math/rand"
	"reflect"
	"testing"

	noteMocks "github.com/ansedo/note-service-api/internal/repository/mocks"
	"github.com/ansedo/note-service-api/internal/service/note"
	desc "github.com/ansedo/note-service-api/pkg/note_v1"

	"github.com/go-faker/faker/v4"
	"github.com/golang/protobuf/ptypes/empty"
	"go.uber.org/mock/gomock"
)

func TestNote_Delete(t *testing.T) {
	var (
		ctx = context.Background()

		noteMock = noteMocks.NewMockNoteRepositoryInterface(gomock.NewController(t))
		api      = NewMockNote(Note{noteService: note.NewMockNoteService(noteMock)})

		validID   = rand.Int63() + 1
		invalidID = int64(0)
		repoErr   = errors.New(faker.Word())
	)
	type args struct {
		ctx context.Context
		req *desc.DeleteRequest
	}
	tests := []struct {
		name    string
		args    args
		repoReq int64
		repoErr error
		want    *empty.Empty
		wantErr bool
	}{
		{
			name: "success case",
			args: args{
				ctx: ctx,
				req: &desc.DeleteRequest{
					Id: validID,
				},
			},
			repoReq: validID,
			repoErr: nil,
			want:    &empty.Empty{},
			wantErr: false,
		},
		{
			name: "repository error",
			args: args{
				ctx: ctx,
				req: &desc.DeleteRequest{
					Id: invalidID,
				},
			},
			repoReq: invalidID,
			repoErr: repoErr,
			want:    nil,
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			noteMock.EXPECT().Delete(tt.args.ctx, tt.repoReq).Return(tt.repoErr)
			got, err := api.Delete(tt.args.ctx, tt.args.req)
			if (err != nil) != tt.wantErr {
				t.Errorf("Delete() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Delete() got = %v, want %v", got, tt.want)
			}
		})
	}
}
