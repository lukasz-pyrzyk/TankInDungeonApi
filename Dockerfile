FROM golang

MAINTAINER Lukasz Pyrzyk <lukasz.pyrzyk@gmail.com>

ADD . /go/src/github.com/lukasz-pyrzyk/BestPlayers

RUN go get gopkg.in/mgo.v2 & go get gopkg.in/yaml.v2 & go get gopkg.in/mgo.v2/internal/json & go get github.com/ant0ine/go-json-rest/rest

RUN go install github.com/lukasz-pyrzyk/BestPlayers/Src

ENTRYPOINT /go/bin/BestPlayers

EXPOSE 8080
