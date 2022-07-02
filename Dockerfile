# syntax=docker/dockerfile:1

FROM golang:1.18-alpine

WORKDIR /app

COPY go.mod ./
COPY *.go ./

RUN go build -o main . 

EXPOSE 80

CMD [ "./main" ]

