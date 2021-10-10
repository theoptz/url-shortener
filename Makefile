test:
	go test ./...

generate-mocks:
	go generate ./internal/interfaces/...

generate-server:
	swagger generate server -f api/api.yaml --exclude-spec --exclude-main -t internal/server

server:
	go build -o cmd/server/server cmd/server/main.go