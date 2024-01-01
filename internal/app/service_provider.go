package app

import (
	"context"
	"log"

	"github.com/ansedo/note-service-api/internal/config"
	"github.com/ansedo/note-service-api/internal/pkg/db"
	"github.com/ansedo/note-service-api/internal/repository"
	"github.com/ansedo/note-service-api/internal/service/note"
)

type serviceProvider struct {
	configPath string
	config     *config.Config

	db *db.Client

	// repositories
	noteRepository *repository.NoteRepository

	// services
	noteService *note.Service
}

func newServiceProvider(pathConfig string) *serviceProvider {
	return &serviceProvider{
		configPath: pathConfig,
	}
}

func (s *serviceProvider) GetDB(ctx context.Context) *db.Client {
	if s.db == nil {
		cfg, err := s.GetConfig().GetDBConfig()
		if err != nil {
			log.Fatalf("failed to get db confg: %s", err.Error())
		}

		dbc, err := db.NewClient(ctx, cfg)
		if err != nil {
			log.Fatalf("failed to connect to db: %s", err.Error())
		}

		s.db = dbc
	}

	return s.db
}

func (s *serviceProvider) GetConfig() *config.Config {
	if s.config == nil {
		cfg, err := config.New(s.configPath)
		if err != nil {
			log.Fatalf("failed to get config: %s", err.Error())
		}

		s.config = cfg
	}

	return s.config
}

func (s *serviceProvider) GetNoteRepository(ctx context.Context) *repository.NoteRepository {
	if s.noteRepository == nil {
		s.noteRepository = repository.NewNoteRepository(s.GetDB(ctx))
	}

	return s.noteRepository
}

func (s *serviceProvider) GetNoteService(ctx context.Context) *note.Service {
	if s.noteService == nil {
		s.noteService = note.NewService(s.GetNoteRepository(ctx))
	}

	return s.noteService
}
