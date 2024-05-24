FROM golang:latest

RUN --mount=type=cache,target=/var/cache/apt \
    apt-get update && apt-get install -y build-essential

WORKDIR /usr/src/app

COPY go.* .

RUN go mod tidy
RUN go mod verify
RUN go mod download

COPY . .