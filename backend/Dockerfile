# syntax=docker/dockerfile:1

FROM golang:1.16-alpine

ADD . /go/src/backend
WORKDIR /go/src/backend

RUN go mod download

COPY *.go ./

RUN go build -o /backend

EXPOSE 8877

CMD [ "/backend" ]