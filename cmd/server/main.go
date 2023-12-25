package main

import (
	"context"
	"log"
	"net"
	"net/http"
	"sync"

	grpcValidator "github.com/grpc-ecosystem/go-grpc-middleware/validator"
	"github.com/grpc-ecosystem/grpc-gateway/runtime"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"

	"github.com/ansedo/note-service-api/internal/app/api/note_v1"
	desc "github.com/ansedo/note-service-api/pkg/note_v1"
)

const (
	hostGrpc = "localhost:50051"
	hostHttp = "localhost:8090"
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

	s := grpc.NewServer(
		grpc.UnaryInterceptor(grpcValidator.UnaryServerInterceptor()),
	)
	desc.RegisterNoteServiceServer(s, note_v1.NewNote())

	log.Printf("grpc server has been started on `%s`", hostGrpc)

	if err = s.Serve(lis); err != nil {
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
