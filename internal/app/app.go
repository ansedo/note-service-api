package app

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

type App struct {
	note            *note_v1.Note
	serviceProvider *serviceProvider

	pathConfig string

	grpcServer *grpc.Server
	mux        *runtime.ServeMux
}

func NewApp(ctx context.Context, pathConfig string) (*App, error) {
	a := &App{
		pathConfig: pathConfig,
	}

	if err := a.initDeps(ctx); err != nil {
		return nil, err
	}

	return a, nil
}

func (a *App) Run() error {
	defer func() {
		a.serviceProvider.db.Close()
	}()

	wg := &sync.WaitGroup{}
	wg.Add(2)

	if err := a.runGRPC(wg); err != nil {
		return err
	}

	if err := a.runPublicHTTP(wg); err != nil {
		return err
	}

	wg.Wait()
	return nil
}

func (a *App) initDeps(ctx context.Context) error {
	inits := []func(ctx2 context.Context) error{
		a.initServiceProvider,
		a.initServer,
		a.initGRPCServer,
		a.initPublicHTTPHandlers,
	}

	for _, f := range inits {
		if err := f(ctx); err != nil {
			return err
		}
	}

	return nil
}

func (a *App) initServiceProvider(ctx context.Context) error {
	a.serviceProvider = newServiceProvider(a.pathConfig)

	return nil
}

func (a *App) initServer(ctx context.Context) error {
	a.note = note_v1.NewNote(a.serviceProvider.GetNoteService(ctx))

	return nil
}

func (a *App) initGRPCServer(ctx context.Context) error {
	a.grpcServer = grpc.NewServer(
		grpc.UnaryInterceptor(grpcValidator.UnaryServerInterceptor()),
	)

	desc.RegisterNoteServiceServer(a.grpcServer, a.note)

	return nil
}

func (a *App) initPublicHTTPHandlers(ctx context.Context) error {
	a.mux = runtime.NewServeMux()

	opts := []grpc.DialOption{grpc.WithTransportCredentials(insecure.NewCredentials())}

	if err := desc.RegisterNoteServiceHandlerFromEndpoint(
		ctx,
		a.mux,
		a.serviceProvider.GetConfig().GRPC.GetAddress(),
		opts,
	); err != nil {
		return err
	}

	return nil
}

func (a *App) runGRPC(wg *sync.WaitGroup) error {
	lis, err := net.Listen("tcp", a.serviceProvider.GetConfig().GRPC.GetAddress())
	if err != nil {
		return err
	}

	go func() {
		defer wg.Done()

		if err = a.grpcServer.Serve(lis); err != nil {
			log.Fatalf("failed to run grpc server: %s", err.Error())
		}
	}()

	log.Printf("grpc server has been started on `%s`", a.serviceProvider.GetConfig().GRPC.GetAddress())
	return nil
}

func (a *App) runPublicHTTP(wg *sync.WaitGroup) error {
	go func() {
		defer wg.Done()

		if err := http.ListenAndServe(a.serviceProvider.GetConfig().HTTP.GetAddress(), a.mux); err != nil {
			log.Fatalf("failed to run http server: %s", err.Error())
		}
	}()

	log.Printf("http server has been started on `%s`", a.serviceProvider.GetConfig().HTTP.GetAddress())
	return nil
}
