PHONY: generate
generate:
	mkdir -p pkg/note_v1
	protoc --proto_path api/note_v1 \
			--go_out=pkg/note_v1 --go_opt=paths=import \
			--go-grpc_out=pkg/note_v1 --go-grpc_opt=paths=import \
			api/note_v1/note.proto
	mv pkg/note_v1/github.com/ansedo/note-service-api/pkg/note_v1/* pkg/note_v1/
	rm -rf pkg/note_v1/github.com
