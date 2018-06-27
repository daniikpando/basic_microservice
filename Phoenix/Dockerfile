FROM golang:1.9

RUN go get github.com/labstack/echo

ADD . /go/src/github.com/daniel/PruebaBackend/Phoenix

WORKDIR /go/src/github.com/daniel/PruebaBackend/Phoenix

RUN go install github.com/daniel/PruebaBackend/Phoenix

ENTRYPOINT /go/bin/Phoenix

EXPOSE 8080
