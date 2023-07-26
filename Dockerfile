FROM golang:1.18-alpine
WORKDIR /app
COPY . /app
RUN go build -o /app main.go
CMD ["/app/main"]
