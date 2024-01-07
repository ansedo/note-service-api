package note_v1

import (
	"context"
	"database/sql"
	"errors"
	"math/rand"
	"reflect"
	"testing"
	"time"

	"github.com/ansedo/note-service-api/internal/model"
	noteMocks "github.com/ansedo/note-service-api/internal/repository/mocks"
	"github.com/ansedo/note-service-api/internal/service/note"
	desc "github.com/ansedo/note-service-api/pkg/note_v1"

	"github.com/go-faker/faker/v4"
	"go.uber.org/mock/gomock"
	"google.golang.org/protobuf/types/known/timestamppb"
)

func TestNote_Get(t *testing.T) {
	var (
		ctx = context.Background()

		noteMock = noteMocks.NewMockNoteRepositoryInterface(gomock.NewController(t))
		api      = NewMockNote(Note{noteService: note.NewMockNoteService(noteMock)})

		validID   = rand.Int63() + 1
		invalidID = int64(0)
		title     = faker.Sentence()
		text      = faker.Paragraph()
		author    = faker.Name()
		email     = faker.Email()
		createdAt = faker.UnixTime()
		repoErr   = errors.New(faker.Word())
	)
	type args struct {
		ctx context.Context
		req *desc.GetRequest
	}
	tests := []struct {
		name    string
		args    args
		repoReq int64
		repoRes *model.Note
		repoErr error
		want    *desc.GetResponse
		wantErr bool
	}{
		{
			name: "success case",
			args: args{
				ctx: ctx,
				req: &desc.GetRequest{
					Id: validID,
				},
			},
			repoReq: validID,
			repoRes: &model.Note{
				ID: validID,
				Info: &model.NoteInfo{
					Title:  title,
					Text:   text,
					Author: author,
					Email:  email,
				},
				CreatedAt: time.Unix(createdAt, 0),
				UpdatedAt: sql.NullTime{},
			},
			repoErr: nil,
			want: &desc.GetResponse{
				Note: &desc.Note{
					Id: validID,
					Info: &desc.NoteInfo{
						Title:  title,
						Text:   text,
						Author: author,
						Email:  email,
					},
					CreatedAt: &timestamppb.Timestamp{
						Seconds: createdAt,
					},
					UpdatedAt: nil,
				},
			},
			wantErr: false,
		},
		{
			name: "repository error",
			args: args{
				ctx: ctx,
				req: &desc.GetRequest{
					Id: invalidID,
				},
			},
			repoReq: invalidID,
			repoRes: nil,
			repoErr: repoErr,
			want:    nil,
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			noteMock.EXPECT().Get(tt.args.ctx, tt.repoReq).Return(tt.repoRes, tt.repoErr)
			got, err := api.Get(tt.args.ctx, tt.args.req)
			if (err != nil) != tt.wantErr {
				t.Errorf("Get() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Get() got = %v, want %v", got, tt.want)
			}
		})
	}
}
