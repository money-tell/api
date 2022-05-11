
deps:
	go mod tidy
	go mod vendor

run:
	go run ./cmd/api.go