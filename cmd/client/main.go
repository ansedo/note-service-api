package main

import (
	"context"
	"log/slog"
	"os"

	desc "github.com/ansedo/note-service-api/pkg/note_v1"
	"github.com/golang/protobuf/ptypes/empty"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

const address = "localhost:50051"

type Client struct {
	svc desc.NoteServiceClient
}

func main() {
	conn, err := grpc.Dial(address, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		slog.Error(
			"failed to connect to grpc server",
			slog.String("op", "cmd.client.main"),
			slog.String("address", address),
			slog.String("error", err.Error()),
		)
		os.Exit(1)
	}
	defer conn.Close()

	c := &Client{
		svc: desc.NewNoteServiceClient(conn),
	}

	ctx := context.Background()

	c.create(ctx)
	c.get(ctx)
	c.getList(ctx)
	c.update(ctx)
	c.delete(ctx)
}

func (c *Client) create(ctx context.Context) {
	const op = "cmd.client.create"

	log := slog.With("op", op)

	req := &desc.CreateRequest{
		Title:  "Title to create",
		Text:   "Text to create",
		Author: "Author to create",
	}

	resp, err := c.svc.Create(ctx, req)
	if err != nil {
		log.Error(
			"failed to create note",
			slog.Any("request", req),
			slog.String("error", err.Error()),
		)
		return
	}

	log.Info(
		"grpc response from `Create` method has been received",
		slog.String("method", "create"),
		slog.Any("request", req),
		slog.Any("response", resp),
	)
}

func (c *Client) get(ctx context.Context) {
	const op = "cmd.client.get"

	log := slog.With("op", op)

	req := &desc.GetRequest{
		Id: 1,
	}

	resp, err := c.svc.Get(ctx, req)
	if err != nil {
		log.Error(
			"failed to get note",
			slog.Any("request", req),
			slog.String("error", err.Error()),
		)
		return
	}

	log.Info(
		"grpc response from `Get` method has been received",
		slog.String("method", "get"),
		slog.Any("request", req),
		slog.Any("response", resp),
	)
}

func (c *Client) getList(ctx context.Context) {
	const op = "cmd.client.getList"

	log := slog.With("op", op)

	req := &empty.Empty{}

	resp, err := c.svc.GetList(ctx, req)
	if err != nil {
		log.Error(
			"failed to get list of notes",
			slog.Any("request", req),
			slog.String("error", err.Error()),
		)
		return
	}

	log.Info(
		"grpc response from `GetList` method has been received",
		slog.String("method", "getList"),
		slog.Any("request", req),
		slog.Any("response", resp),
	)
}

func (c *Client) update(ctx context.Context) {
	const op = "cmd.client.update"

	log := slog.With("op", op)

	req := &desc.UpdateRequest{
		Note: &desc.Note{
			Id:     1,
			Title:  "Title to update",
			Text:   "Text to update",
			Author: "Author to update",
		},
	}

	resp, err := c.svc.Update(ctx, req)
	if err != nil {
		log.Error(
			"failed to update note",
			slog.Any("request", req),
			slog.String("error", err.Error()),
		)
		return
	}

	log.Info(
		"grpc response from `Update` method has been received",
		slog.String("method", "update"),
		slog.Any("request", req),
		slog.Any("response", resp),
	)
}

func (c *Client) delete(ctx context.Context) {
	const op = "cmd.client.delete"

	log := slog.With("op", op)

	req := &desc.DeleteRequest{
		Id: 1,
	}

	resp, err := c.svc.Delete(ctx, req)
	if err != nil {
		log.Error(
			"failed to delete note",
			slog.Any("request", req),
			slog.String("error", err.Error()),
		)
		return
	}

	log.Info(
		"grpc response from `Delete` method has been received",
		slog.String("method", "delete"),
		slog.Any("request", req),
		slog.Any("response", resp),
	)
}
