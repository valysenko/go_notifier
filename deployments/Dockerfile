FROM golang:latest as builder
WORKDIR /go/src/
COPY . .
RUN go mod download
RUN go mod verify
RUN CGO_ENABLED=0 GOOS=linux go build -ldflags "-s -w" -o bin/go_notifier ./cmd/main.go

FROM alpine
COPY --from=builder /go/src/bin/* /
CMD [ "/go_notifier" ]