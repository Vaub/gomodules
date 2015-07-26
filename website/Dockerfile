FROM alpine:3.2
MAINTAINER Vincent <vincent.aube@hotmail.com>

RUN apk add --update git && apk add --update go
RUN mkdir -p /goenv/src /goenv/pkg /goenv/bin

ENV GOPATH=/goenv

RUN go get github.com/vaub/gomodules/
RUN go get github.com/gorilla/mux

WORKDIR /goenv/src/github.com/vaub/gomodules/website
CMD ["go", "run", "website.go"]
