FROM golang

MAINTAINER Lukasz Pyrzyk <lukasz.pyrzyk@gmail.com>

ADD . /go/src/github.com/lukasz-pyrzyk/BestPlayers

RUN go install github.com/lukasz-pyrzyk/BestPlayers

ENTRYPOINT /go/bin/BestPlayers

EXPOSE 8080