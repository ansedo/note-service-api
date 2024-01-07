package note_v1

import (
	"context"
	"errors"
	"math/rand"
	"reflect"
	"testing"

	"github.com/ansedo/note-service-api/internal/model"
	noteMocks "github.com/ansedo/note-service-api/internal/repository/mocks"
	"github.com/ansedo/note-service-api/internal/service/note"
	desc "github.com/ansedo/note-service-api/pkg/note_v1"

	"github.com/go-faker/faker/v4"
	"go.uber.org/mock/gomock"
)

func TestNote_Create(t *testing.T) {
	var (
		ctx = context.Background()

		noteMock = noteMocks.NewMockNoteRepositoryInterface(gomock.NewController(t))
		api      = NewMockNote(Note{noteService: note.NewMockNoteService(noteMock)})

		validID      = rand.Int63() + 1
		invalidID    = int64(0)
		validTitle   = faker.Sentence()
		validText    = faker.Paragraph()
		validAuthor  = faker.Name()
		validEmail   = faker.Email()
		invalidEmail = faker.Word()
		repoErr      = errors.New(faker.Word())
	)
	type args struct {
		ctx context.Context
		req *desc.CreateRequest
	}
	tests := []struct {
		name    string
		args    args
		repoReq *model.NoteInfo
		repoRes int64
		repoErr error
		want    *desc.CreateResponse
		wantErr bool
	}{
		{
			name: "success case",
			args: args{
				ctx: ctx,
				req: &desc.CreateRequest{
					Note: &desc.NoteInfo{
						Title:  validTitle,
						Text:   validText,
						Author: validAuthor,
						Email:  validEmail,
					},
				},
			},
			repoReq: &model.NoteInfo{
				Title:  validTitle,
				Text:   validText,
				Author: validAuthor,
				Email:  validEmail,
			},
			repoRes: validID,
			repoErr: nil,
			want:    &desc.CreateResponse{Id: validID},
			wantErr: false,
		},
		{
			name: "repository error",
			args: args{
				ctx: ctx,
				req: &desc.CreateRequest{
					Note: &desc.NoteInfo{
						Title:  validTitle,
						Text:   validText,
						Author: validAuthor,
						Email:  invalidEmail,
					},
				},
			},
			repoReq: &model.NoteInfo{
				Title:  validTitle,
				Text:   validText,
				Author: validAuthor,
				Email:  invalidEmail,
			},
			repoRes: invalidID,
			repoErr: repoErr,
			want:    nil,
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			noteMock.EXPECT().Create(tt.args.ctx, tt.repoReq).Return(tt.repoRes, tt.repoErr)
			got, err := api.Create(tt.args.ctx, tt.args.req)
			if (err != nil) != tt.wantErr {
				t.Errorf("Create() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Create() got = %v, want %v", got, tt.want)
			}
		})
	}
}
