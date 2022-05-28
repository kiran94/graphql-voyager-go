build:
	go build ./...

test:
	gotestsum --format testname -- -race ./...

install_tools:
	go install gotest.tools/gotestsum@latest

run_server:
	go run ./cmd/server
