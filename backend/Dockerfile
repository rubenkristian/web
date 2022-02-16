FROM golang:1.17.7-alpine3.15 AS builder

RUN apk --no-cache add gcc g++ make
RUN apk add git
WORKDIR /go/app

COPY go.mod go.sum ./
RUN go mod download && go mod verify

COPY . .
RUN GOOS=linux GOARCH=amd64 go build -ldflags="-w -s" -o /go/bin/app

FROM scratch

COPY --from=builder /go/bin/app /go/bin/app

EXPOSE 8080
ENTRYPOINT ["./go/bin/app"]