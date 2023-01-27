FROM golang:1.19.1-alpine3.16 AS build_base

ENV CGO_ENABLED=1
ENV GO111MODULE=on
RUN apk add --no-cache git  git gcc g++

# Set the Current Working Directory inside the container
WORKDIR /src

# We want to populate the module cache based on the go.{mod,sum} files.
COPY go.mod .
COPY go.sum .

RUN go mod download

COPY . .

# Build the Go app
RUN go build -o ./out/app ./main.go

# Start fresh from a smaller image
FROM alpine:3.12
RUN apk add ca-certificates

WORKDIR /app

COPY --from=build_base /src/out/app /app/restapi
COPY --from=build_base /src/data /app/data

RUN chmod +x restapi

# This container exposes port 3000 to the outside world
EXPOSE 3000

# Run the binary program produced by `go install`
ENTRYPOINT ./restapi