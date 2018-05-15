FROM golang:1.10

ADD . /go/src/github.com/Sharykhin/go-users-api

WORKDIR /go/src/github.com/Sharykhin/go-users-api

RUN curl https://raw.githubusercontent.com/golang/dep/master/install.sh | sh

RUN dep ensure && \
    go build -o /go/bin/go-users-api .

EXPOSE 8080

ENTRYPOINT /go/bin/go-users-api