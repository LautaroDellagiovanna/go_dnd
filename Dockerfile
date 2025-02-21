# syntax=docker/dockerfile:1

FROM golang:1.21

WORKDIR /go_dnd_docker

COPY go.mod go.sum ./

RUN go mod download

COPY . .

RUN go build -o main ./cmd/api/main.go

EXPOSE 9090
EXPOSE 8080
CMD ["./main"]

