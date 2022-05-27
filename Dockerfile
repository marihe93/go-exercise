# syntax=docker/dockerfile:1

# Using alpine due to its smaller footprint
FROM golang:1.14-alpine

WORKDIR /app

# Downloading required Go modules
COPY go.mod ./
COPY go.sum ./
RUN go mod download

# Copyying all go files
COPY *.go ./

# Building application
RUN go build -o /go-exercise

# Exposing port 8080 for API access
EXPOSE 8080

CMD [ "/go-exercise" ]
