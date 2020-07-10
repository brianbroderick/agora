FROM golang:1.13-alpine as builder
RUN apk update && apk upgrade && \
  apk add --no-cache git openssh

WORKDIR /go/src/github.com/brianbroderick/agora
COPY go.mod .
COPY go.sum .

RUN go mod download

COPY . . 

CMD ["go test"]