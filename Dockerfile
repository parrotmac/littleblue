FROM golang:latest as builder

# Install script isn't cross-platform
RUN go get -u github.com/golang/dep/cmd/dep

RUN mkdir -p /go/src/github.com/parrotmac/littleblue/
WORKDIR /go/src/github.com/parrotmac/littleblue/
COPY . .

RUN dep ensure -v -vendor-only
RUN go build -o bin/littleblue cmd/main.go

FROM alpine:latest

RUN mkdir -p /opt/littleblue
COPY --from=builder /go/src/github.com/parrotmac/littleblue/bin/littleblue /opt/littleblue

CMD ["/opt/littleblue/littleblue"]
