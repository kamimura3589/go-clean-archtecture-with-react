FROM golang:1.11.5-alpine3.9

RUN apk update && apk upgrade && \
    apk --update add git gcc make && \
    go get -u github.com/golang/dep/cmd/dep

WORKDIR /go/src/github.com/kamimura3589/clean-architecture-go

COPY . .

RUN make install 
