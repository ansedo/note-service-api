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
	"github.com/golang/protobuf/ptypes/empty"
	"go.uber.org/mock/gomock"
	"google.golang.org/protobuf/types/known/timestamppb"
)

func TestNote_GetList(t *testing.T) {
	var (
		ctx = context.Background()

		noteMock = noteMocks.NewMockNoteRepositoryInterface(gomock.NewController(t))
		api      = NewMockNote(Note{noteService: note.NewMockNoteService(noteMock)})

		ids        = []int64{rand.Int63() + 1, rand.Int63() + 1, rand.Int63() + 1}
		titles     = []string{faker.Sentence(), faker.Sentence(), faker.Sentence()}
		texts      = []string{faker.Paragraph(), faker.Paragraph(), faker.Paragraph()}
		authors    = []string{faker.Name(), faker.Name(), faker.Name()}
		emails     = []string{faker.Email(), faker.Email(), faker.Email()}
		createdAts = []int64{faker.UnixTime(), faker.UnixTime(), faker.UnixTime()}
		updatedAts = []int64{faker.UnixTime(), faker.UnixTime(), faker.UnixTime()}
		repoErr    = errors.New(faker.Word())
	)
	type args struct {
		ctx context.Context
		req *empty.Empty
	}
	repoRes := make([]*model.Note, len(ids))
	for i := range ids {
		repoRes[i] = &model.Note{
			ID: ids[i],
			Info: &model.NoteInfo{
				Title:  titles[i],
				Text:   texts[i],
				Author: authors[i],
				Email:  emails[i],
			},
			CreatedAt: time.Unix(createdAts[i], 0),
			UpdatedAt: sql.NullTime{
				Time:  time.Unix(updatedAts[i], 0),
				Valid: true,
			},
		}
	}
	want := &desc.GetListResponse{
		Notes: make([]*desc.Note, len(ids)),
	}
	for i := range ids {
		want.Notes[i] = &desc.Note{
			Id: ids[i],
			Info: &desc.NoteInfo{
				Title:  titles[i],
				Text:   texts[i],
				Author: authors[i],
				Email:  emails[i],
			},
			CreatedAt: &timestamppb.Timestamp{
				Seconds: createdAts[i],
			},
			UpdatedAt: &timestamppb.Timestamp{
				Seconds: updatedAts[i],
			},
		}
	}
	tests := []struct {
		name    string
		args    args
		repoRes []*model.Note
		repoErr error
		want    *desc.GetListResponse
		wantErr bool
	}{
		{
			name: "success case",
			args: args{
				ctx: ctx,
				req: &empty.Empty{},
			},
			repoRes: repoRes,
			repoErr: nil,
			want:    want,
			wantErr: false,
		},
		{
			name: "repository error",
			args: args{
				ctx: ctx,
				req: &empty.Empty{},
			},
			repoRes: nil,
			repoErr: repoErr,
			want:    nil,
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			noteMock.EXPECT().GetList(tt.args.ctx).Return(tt.repoRes, tt.repoErr)
			got, err := api.GetList(tt.args.ctx, tt.args.req)
			if (err != nil) != tt.wantErr {
				t.Errorf("GetList() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("GetList() got = %v, want %v", got, tt.want)
			}
		})
	}
}
