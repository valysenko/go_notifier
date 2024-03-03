FROM golang:latest as builder
WORKDIR /go/src/
COPY . .
RUN go mod download
RUN go mod verify
CMD CGO_ENABLED=0 go test ./...