# syntax=docker/dockerfile:1
FROM golang:1.16-alpine

# setup a work directory for go
WORKDIR /app

# copy go files to workdir for access
COPY go.mod ./
COPY go.sum ./

# run go in the image
RUN go mod download

# copy source code to image
COPY *.go ./

# compile the app
RUN go build -o /twitter-clone

# use port number below
EXPOSE 8080

# execute app
CMD [ "/twitter-clone" ]