FROM node:alpine as frontend

# node-sass may not provide a prebuilt (e.g. for arm)
RUN apk update
RUN apk add python make g++

RUN mkdir -p /app/client
WORKDIR /app/client

COPY client/package.json .
COPY client/tsconfig.json .
COPY client/yarn.lock .

RUN yarn

COPY client/ .

RUN yarn build
# Final artifact: /app/client/build

FROM golang:latest as builder

ENV CGO_ENABLED=0

WORKDIR /app
COPY . .

RUN go build -o bin/littleblue cmd/main.go
# Final artifact: /app/bin/littleblue

FROM alpine:latest

# Ensure runtime has valid certs
RUN apk update \
        && apk upgrade \
        && apk add --no-cache \
        ca-certificates \
        && update-ca-certificates 2>/dev/null || true

RUN mkdir -p /opt/littleblue
WORKDIR /opt/littleblue
COPY --from=builder /app/bin/littleblue .
COPY --from=frontend /app/client/build ./client/build
COPY config.yml .

CMD ["/opt/littleblue/littleblue"]
