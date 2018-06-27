FROM golang:1.9

RUN go get github.com/labstack/echo
RUN go get github.com/lib/pq
RUN go get github.com/jinzhu/gorm

ADD . /go/src/github.com/daniel/PruebaBackend/UserMicroservice

WORKDIR /go/src/github.com/daniel/PruebaBackend/UserMicroservice

RUN go install github.com/daniel/PruebaBackend/UserMicroservice

ENTRYPOINT /go/bin/UserMicroservice

EXPOSE 8080
