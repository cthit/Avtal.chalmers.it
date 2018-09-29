FROM golang:1.10
MAINTAINER digIT <digit@chalmers.it>

RUN mkdir -p /go/src/github.com/avtal.chalmers.it/backend
WORKDIR /go/src/github.com/avtal.chalmers.it/backend

COPY /backend /go/src/github.com/avtal.chalmers.it/backend

RUN go get -d
RUN go install
CMD go run main.go

EXPOSE 3000
