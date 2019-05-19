FROM node:alpine as frontend


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

# Install script isn't cross-platform
RUN go get -u github.com/golang/dep/cmd/dep

RUN mkdir -p /go/src/github.com/parrotmac/littleblue/
WORKDIR /go/src/github.com/parrotmac/littleblue/
COPY Gopkg.toml .
COPY Gopkg.lock .

RUN dep ensure -v -vendor-only
COPY . .
RUN go build -o bin/littleblue cmd/main.go
# Final artifact: /go/src/github.com/parrotmac/littleblue/bin/littleblue

FROM alpine:latest

RUN mkdir -p /opt/littleblue
WORKDIR /opt/littleblue
COPY --from=builder /go/src/github.com/parrotmac/littleblue/bin/littleblue .
COPY --from=frontend /app/client/build ./client/build
COPY config.yml .

CMD ["/opt/littleblue/littleblue"]
