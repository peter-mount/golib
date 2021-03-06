# Dockerfile used to test compile the libraries

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

WORKDIR /work
ADD . .
RUN go build ./...
