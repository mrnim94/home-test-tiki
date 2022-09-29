FROM golang:1.14.6-alpine

RUN apk update && apk add git

ENV CGO_ENABLED=0
ENV GO111MODULE=on

ENV GOPATH /go
ENV PATH $GOPATH/bin:/usr/local/go/bin:$PATH

RUN mkdir -p "$GOPATH/src" "$GOPATH/bin" && chmod -R 777 "$GOPATH"
WORKDIR $GOPATH/src/home-test-tiki

COPY . .

# RUN go mod init home-test-tiki
WORKDIR cmd/production
RUN GOOS=linux go build -mod vendor -o app

ENTRYPOINT ["./app"]

EXPOSE 80
