FROM golang:1.10
MAINTAINER digIT <digit@chalmers.it>

RUN mkdir -p /go/src/github.com/avtal.chalmers.it/frontend
WORKDIR /go/src/github.com/avtal.chalmers.it/frontend

COPY /frontend /go/src/github.com/avtal.chalmers.it/frontend

RUN go get -d
RUN go install
CMD go run main.go

EXPOSE 8080
