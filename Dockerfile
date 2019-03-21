FROM golang:latest

RUN mkdir -p /go/src/timezones
WORKDIR /go/src/timezones

ADD . /go/src/timezones

RUN go get -v
RUN go get github.com/stretchr/testify/assert