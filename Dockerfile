FROM golang:1.8.3

ENV GOPATH /go

VOLUME ["/uploads"]
VOLUME ["/data"]

RUN mkdir /app
RUN mkdir -p /go/src/volcanofm

WORKDIR /go/src/volcanofm
COPY . .

RUN go-wrapper download
RUN go-wrapper install