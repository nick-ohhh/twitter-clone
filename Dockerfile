# syntax=docker/dockerfile:1
FROM golang:1.16-alpine

# setup a work directory for go
WORKDIR /app

# copy go files to workdir for access
COPY go.mod ./
COPY go.sum ./

# run go in the image
RUN go mod download

#