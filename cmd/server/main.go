package main

import (
	"log/slog"
	"net"
	"os"

	"github.com/ansedo/note-service-api/internal/app/api/note_v1"

	desc "github.com/ansedo/note-service-api/pkg/note_v1"
	"google.golang.org/grpc"
)

const port = ":50051"

func main() {
	const op = "cmd.server.main"

	log := slog.With("op", op)

	lis, err := net.Listen("tcp", port)
	if err != nil {
		log.Error(
			"failed to mapping port",
			slog.String("port", port),
			slog.String("error", err.Error()),
		)
		os.Exit(1)
	}

	s := grpc.NewServer()
	desc.RegisterNoteServiceServer(s, note_v1.NewNote())

	log.Info(
		"grpc server has been started",
		slog.String("port", port),
	)

	if err = s.Serve(lis); err != nil {
		log.Error(
			"failed to serve grpc",
			slog.Any("listener", lis),
			slog.String("error", err.Error()),
		)
		os.Exit(1)
	}
}
