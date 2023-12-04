package note_v1

import (
	"fmt"

	desc "github.com/ansedo/note-service-api/pkg/note_v1"
)

const (
	host       = "localhost"
	port       = "54321"
	dbUser     = "note-service-user"
	dbPassword = "note-service-password"
	dbName     = "note-service"
	sslMode    = "disable"

	noteTable          = "note"
	sqlColumnId        = "id"
	sqlColumnTitle     = "title"
	sqlColumnText      = "text"
	sqlColumnAuthor    = "author"
	sqlColumnUpdatedAt = "updated_at"
)

var dbDsn = fmt.Sprintf(
	"host=%s port=%s user=%s password=%s dbname=%s sslmode=%s",
	host, port, dbUser, dbPassword, dbName, sslMode,
)

type Note struct {
	desc.UnimplementedNoteServiceServer
}

func NewNote() *Note {
	return &Note{}
}
