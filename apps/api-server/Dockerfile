FROM golang:1.24.3 AS builder

WORKDIR /app
COPY . .


RUN go mod tidy
RUN GOOS=linux GOARCH=amd64 go build -o server ./cmd/server

FROM debian:bullseye-slim
WORKDIR /app

COPY --from=builder /app/server .

EXPOSE 8080

CMD ["./server"]