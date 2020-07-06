FROM golang:alpine AS builder

RUN apk update && apk add --no-cache git

WORKDIR $GOPATH/

COPY . .

RUN go get -d -v

RUN go build -o /go/bin/sum
