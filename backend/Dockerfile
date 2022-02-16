FROM golang:1.17.7-alpine3.15 AS builder

RUN apk --no-cache add gcc g++ make
RUN apk add git
WORKDIR /go/src/app

COPY go.mod go.sum ./
RUN go mod download && go mod verify

COPY . .
RUN GOOS=linux GOARCH=amd64 go build -ldflags="-w -s" -o ./bin/server ./app.go

FROM alpine:3.10
RUN apk --no-cache add ca-certificates
WORKDIR /usr/bin
COPY --from=builder /go/src/app/bin /go/bin
EXPOSE 8080
ENTRYPOINT /go/bin/server --port 8080
