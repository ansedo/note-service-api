package note_v1

import (
	"context"
	"database/sql"
	"errors"
	"math/rand"
	"reflect"
	"testing"

	"github.com/ansedo/note-service-api/internal/model"
	noteMocks "github.com/ansedo/note-service-api/internal/repository/mocks"
	"github.com/ansedo/note-service-api/internal/service/note"
	desc "github.com/ansedo/note-service-api/pkg/note_v1"

	"github.com/go-faker/faker/v4"
	"github.com/golang/protobuf/ptypes/empty"
	"go.uber.org/mock/gomock"
	"google.golang.org/protobuf/types/known/wrapperspb"
)

func TestNote_Update(t *testing.T) {
	var (
		ctx = context.Background()

		noteMock = noteMocks.NewMockNoteRepositoryInterface(gomock.NewController(t))
		api      = NewMockNote(Note{noteService: note.NewMockNoteService(noteMock)})

		validID      = rand.Int63() + 1
		validTitle   = faker.Sentence()
		validEmail   = faker.Email()
		invalidEmail = faker.Word()
		repoErr      = errors.New(faker.Word())
	)
	type args struct {
		ctx context.Context
		req *desc.UpdateRequest
	}
	type repoReq struct {
		id             int64
		updateNoteInfo *model.UpdateNoteInfo
	}
	tests := []struct {
		name    string
		args    args
		repoReq repoReq
		repoErr error
		want    *empty.Empty
		wantErr bool
	}{
		{
			name: "success case",
			args: args{
				ctx: ctx,
				req: &desc.UpdateRequest{
					Id: validID,
					Note: &desc.UpdateNoteInfo{
						Title: &wrapperspb.StringValue{Value: validTitle},
						Email: &wrapperspb.StringValue{Value: validEmail},
					},
				},
			},
			repoReq: repoReq{
				id: validID,
				updateNoteInfo: &model.UpdateNoteInfo{
					Title: sql.NullString{
						String: validTitle,
						Valid:  true,
					},
					Email: sql.NullString{
						String: validEmail,
						Valid:  true,
					},
				},
			},
			repoErr: nil,
			want:    &empty.Empty{},
			wantErr: false,
		},
		{
			name: "repository error",
			args: args{
				ctx: ctx,
				req: &desc.UpdateRequest{
					Id: validID,
					Note: &desc.UpdateNoteInfo{
						Title: &wrapperspb.StringValue{Value: validTitle},
						Email: &wrapperspb.StringValue{Value: invalidEmail},
					},
				},
			},
			repoReq: repoReq{
				id: validID,
				updateNoteInfo: &model.UpdateNoteInfo{
					Title: sql.NullString{
						String: validTitle,
						Valid:  true,
					},
					Email: sql.NullString{
						String: invalidEmail,
						Valid:  true,
					},
				},
			},
			repoErr: repoErr,
			want:    nil,
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			noteMock.EXPECT().Update(tt.args.ctx, tt.repoReq.id, tt.repoReq.updateNoteInfo).Return(tt.repoErr)
			got, err := api.Update(tt.args.ctx, tt.args.req)
			if (err != nil) != tt.wantErr {
				t.Errorf("Update() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Update() got = %v, want %v", got, tt.want)
			}
		})
	}
}
