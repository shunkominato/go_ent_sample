FROM golang:1.21.3-alpine

RUN apk update && apk add git && apk add alpine-sdk && apk add postgresql-client && apk add zsh

ENV APP_ROOT /go/src/app

ENV SHELL /usr/bin/zsh

RUN mkdir $APP_ROOT

WORKDIR $APP_ROOT

ENV CGO_ENABLED=1 \
  GOOS=linux \
  GOARCH=arm64 \
  GO111MODULE=on

COPY . $APP_ROOT

# Air: hot reload
RUN go install github.com/cosmtrek/air@latest
RUN go install github.com/go-delve/delve/cmd/dlv@latest

RUN curl -sSf https://atlasgo.sh | sh

EXPOSE 8080