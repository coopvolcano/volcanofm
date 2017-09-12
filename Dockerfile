FROM golang:1.8.3

ENV GOPATH /go

RUN mkdir /music
RUN mkdir /app
RUN mkdir -p /go/src/volcanofm

ADD . /go/src/volcanofm

WORKDIR /go/src/volcanofm

RUN go get -v
RUN go build -o /app/volcanofm .