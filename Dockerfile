# Dockerfile used to test compile the libraries

ARG moduleName

# ============================================================
# Build container containing our pre-pulled libraries.
# As this changes rarely it means we can use the cache between
# building each microservice.
FROM golang:alpine as build

# The golang alpine image is missing git so ensure we have additional tools
RUN apk add --no-cache \
      curl \
      git

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

FROM build as compiler
ARG moduleName

RUN go build \
      -v \
      github.com/peter-mount/golib/${moduleName}
