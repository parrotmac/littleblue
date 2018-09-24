FROM golang:latest as gobuilder

WORKDIR /go/src/app
COPY . .

RUN go get -d -v ./
RUN go build -v ./

RUN mkdir /go/src/app/workdir

ENTRYPOINT ["/go/src/app/app"]

#RUN ls
#
#FROM docker:git
#
#RUN mkdir -p /app
#WORKDIR /app
#RUN mkdir workdir
#
#COPY --from=gobuilder /go/src/app/app .
#
#RUN stat /app/app
#RUN ls
#RUN pwd
#
#ENTRYPOINT ["/bin/sh"]
