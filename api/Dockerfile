# Dockerfile.protogen
FROM golang:latest

LABEL maintainer="Will Fanton <willf@ntom.dev>"

ENV PROTOC_VERSION 3.6.1
ENV PROTOC_GEN_GO_VERSION v1.2.0

WORKDIR /go/src/github.com/willfantom/go-iperf/api

RUN apt update
RUN apt install -y protobuf-compiler python3 python3-pip
RUN go get -u github.com/golang/protobuf/protoc-gen-go
RUN pip3 install grpcio-tools
RUN export PATH=$PATH:$GOPATH/bin
RUN echo $PATH

COPY proto/control.proto proto/control.proto

RUN mkdir go
RUN protoc -I proto/. --go_out=plugins=grpc:go control.proto
