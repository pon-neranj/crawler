ARG BUILDER_IMAGE="golang:1.20"
ARG RUNNER_IMAGE="alpine:3.18.4"

FROM ${BUILDER_IMAGE} AS builder

WORKDIR /app

ADD . /app/

RUN go mod tidy
RUN go mod verify

WORKDIR /app/cmd/api

RUN go build -v

CMD ["go","run","."]