FROM golang:1.16-alpine
WORKDIR /app
COPY . /app
RUN go build -o /app  cmd/server/main.go
CMD ["/app/main"]
