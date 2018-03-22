FROM golang:alpine3.7 AS build

RUN apk --update add git binutils
RUN go get -v github.com/golang/dep && \
    go install -v github.com/golang/dep/cmd/dep

ADD . /go/src/github.com/tei1988/iap-grant-role

RUN cd /go/src/github.com/tei1988/iap-grant-role && \
    dep ensure && \
    go build && \
    strip iap-grant-role

FROM alpine:3.7

WORKDIR /opt/app

COPY --from=build /go/src/github.com/tei1988/iap-grant-role/iap-grant-role .
