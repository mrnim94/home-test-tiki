FROM golang:1.14.6-alpine

RUN apk update && apk add git

ENV CGO_ENABLED=0
ENV GO111MODULE=on

ENV GOPATH /go
ENV PATH $GOPATH/bin:/usr/local/go/bin:$PATH

RUN mkdir -p "$GOPATH/src" "$GOPATH/bin" && chmod -R 777 "$GOPATH"
WORKDIR $GOPATH/src/home-test-tiki

COPY . .
RUN go mod init home-test-tiki
RUN go get -v github.com/labstack/echo/v4
RUN go get -v github.com/lestrrat/go-file-rotatelogs
RUN go get -v github.com/rifflock/lfshook
RUN go get -v github.com/sirupsen/logrus
WORKDIR cmd/production
RUN GOOS=linux go build -o app

# ENTRYPOINT ["./app"]

# EXPOSE 80
