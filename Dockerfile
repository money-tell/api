FROM golang:1.19

RUN go install github.com/kyleconroy/sqlc/cmd/sqlc@latest && \
    go install github.com/pressly/goose/v3/cmd/goose@latest && \
    go install github.com/vektra/mockery/v2@latest
