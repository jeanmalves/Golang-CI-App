FROM golang:1.14

COPY src/ /go/src/

WORKDIR /go/src/github.com/jeanmalves/operation/app

RUN go install

ENTRYPOINT ["/go/bin/app"]