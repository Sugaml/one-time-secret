# syntax=docker/dockerfile:experimental
FROM golangci/golangci-lint:v1.50.0 AS base

WORKDIR /app
COPY go.* ./
RUN --mount=type=ssh go mod download
COPY . .
RUN --mount=type=ssh go mod tidy

FROM base as lint
RUN golangci-lint run --timeout 10m0s ./...

FROM base as test
RUN go test -v -coverprofile=cover.out ./...
RUN go tool cover -func=cover.out

FROM base as builder
RUN --mount=type=ssh CGO_ENABLED=0 GOOS=linux go build -o main .

FROM alpine:latest
WORKDIR /app
RUN touch .env
COPY --from=builder /app/main .
CMD ["./main"]
