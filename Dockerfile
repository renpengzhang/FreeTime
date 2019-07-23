FROM golang:alpine

RUN mkdir /go/src/FreeTime

ADD . /go/src/FreeTime

WORKDIR /go/src/FreeTime

RUN apk add --no-cache git
RUN go get -d -v ./...
RUN go install -v ./...

RUN go build -o main .

RUN apk --no-cache add ca-certificates

ENTRYPOINT ./main

EXPOSE 8080