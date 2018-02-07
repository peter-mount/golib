# Dockerfile used to test compile the libraries

# Build container containing our pre-pulled libraries
FROM golang:latest as build

# Static compile
ENV CGO_ENABLED=0
ENV GOOS=linux

# Ensure we have the libraries - docker will cache these between builds
RUN go get -v \
      flag \
      github.com/gorilla/handlers \
      github.com/gorilla/mux \
      github.com/streadway/amqp \
      gopkg.in/robfig/cron.v2 \
      gopkg.in/yaml.v2 \
      io/ioutil \
      log \
      net/http \
      path/filepath \
      time

# Import the source and compile
WORKDIR /go/src/github.com/peter-mount/golib
ADD . /go/src/github.com/peter-mount/golib

RUN go build \
      -v \
      github.com/peter-mount/golib/codec

RUN go build \
      -v \
      github.com/peter-mount/golib/rabbitmq

RUN go build \
      -v \
      github.com/peter-mount/golib/rest

RUN go build \
      -v \
      github.com/peter-mount/golib/statistics

RUN go build \
      -v \
      github.com/peter-mount/golib/util
