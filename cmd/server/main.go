package main

import (
	"context"
	"fmt"
	"log"
	"net"
	"net/http"
	"sync"

	grpcValidator "github.com/grpc-ecosystem/go-grpc-middleware/validator"
	"github.com/grpc-ecosystem/grpc-gateway/runtime"
	_ "github.com/jackc/pgx/stdlib"
	"github.com/jmoiron/sqlx"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"

	"github.com/ansedo/note-service-api/internal/app/api/note_v1"
	"github.com/ansedo/note-service-api/internal/repository"
	"github.com/ansedo/note-service-api/internal/service/note"
	desc "github.com/ansedo/note-service-api/pkg/note_v1"
)

const (
	hostGrpc = "localhost:50051"
	hostHttp = "localhost:8090"

	host       = "localhost"
	port       = "54321"
	dbUser     = "note-service-user"
	dbPassword = "note-service-password"
	dbName     = "note-service"
	sslMode    = "disable"
)

var dbDsn = fmt.Sprintf(
	"host=%s port=%s user=%s password=%s dbname=%s sslmode=%s",
	host, port, dbUser, dbPassword, dbName, sslMode,
)

func main() {
	wg := sync.WaitGroup{}
	wg.Add(2)

	go func() {
		defer wg.Done()
		startGRPC()
	}()
	go func() {
		defer wg.Done()
		startHttp()
	}()

	wg.Wait()
}

func startGRPC() error {
	lis, err := net.Listen("tcp", hostGrpc)
	if err != nil {
		log.Fatalf("failed to mapping hostGrpc `%s`: %s", hostGrpc, err.Error())
	}

	db, err := sqlx.Open("pgx", dbDsn)
	if err != nil {
		return err
	}
	defer db.Close()

	noteService := note.NewService(
		repository.NewNoteRepository(db),
	)

	srv := grpc.NewServer(
		grpc.UnaryInterceptor(grpcValidator.UnaryServerInterceptor()),
	)
	desc.RegisterNoteServiceServer(srv, note_v1.NewNote(noteService))

	log.Printf("grpc server has been started on `%s`", hostGrpc)

	if err = srv.Serve(lis); err != nil {
		log.Fatalf("failed to serve grpc: %s", err.Error())
	}

	return nil
}

func startHttp() error {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	mux := runtime.NewServeMux()
	opts := []grpc.DialOption{grpc.WithTransportCredentials(insecure.NewCredentials())}
	err := desc.RegisterNoteServiceHandlerFromEndpoint(ctx, mux, hostGrpc, opts)
	if err != nil {
		return err
	}

	log.Printf("http server has been started on `%s`", hostHttp)

	return http.ListenAndServe(hostHttp, mux)
}
