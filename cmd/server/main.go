package main

import (
	"log"
	"net"

	"github.com/ansedo/note-service-api/internal/app/api/note_v1"

	desc "github.com/ansedo/note-service-api/pkg/note_v1"
	"google.golang.org/grpc"
)

const port = ":50051"

func main() {
	lis, err := net.Listen("tcp", port)
	if err != nil {
		log.Fatalf("failed to mapping port `%s`: %s", port, err.Error())
	}

	s := grpc.NewServer()
	desc.RegisterNoteServiceServer(s, note_v1.NewNote())

	log.Printf("grpc server has been started on port `%s`", port)

	if err = s.Serve(lis); err != nil {
		log.Fatalf("failed to serve grpc: %s", err.Error())
	}
}
