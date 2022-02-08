FROM golang:1.17

WORKDIR /go/src/app

ENV PROTOC_ZIP=protoc-3.19.3-linux-x86_64.zip
RUN apt-get update && apt-get install -y unzip
RUN curl -OL https://github.com/protocolbuffers/protobuf/releases/download/v3.19.3/$PROTOC_ZIP \
    && unzip -o $PROTOC_ZIP -d /usr/local bin/protoc \
    && unzip -o $PROTOC_ZIP -d /usr/local 'include/*' \ 
    && rm -f $PROTOC_ZIP

RUN go get -u github.com/golang/protobuf/protoc-gen-go

RUN go get github.com/silenceper/gowatch
