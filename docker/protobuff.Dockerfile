FROM golang:1.24

RUN apt-get update && apt-get install -y protobuf-compiler

RUN git clone https://github.com/googleapis/googleapis.git /googleapis

RUN go install google.golang.org/protobuf/cmd/protoc-gen-go@latest
RUN go install google.golang.org/grpc/cmd/protoc-gen-go-grpc@latest
RUN go install github.com/grpc-ecosystem/grpc-gateway/v2/protoc-gen-grpc-gateway@latest

ENV PATH="$PATH:/go/bin"

WORKDIR /app