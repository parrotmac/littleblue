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

FROM alpine:latest

RUN mkdir -p /opt/littleblue
WORKDIR /opt/littleblue
COPY --from=builder /go/src/github.com/parrotmac/littleblue/bin/littleblue /opt/littleblue
COPY config.yml /opt/littleblue

CMD ["/opt/littleblue/littleblue"]
