# syntax=docker/dockerfile:1

FROM golang:1.18-alpine

WORKDIR /app

COPY *.go ./

RUN go build main.go 

EXPOSE 80

CMD [ "./main" ]

