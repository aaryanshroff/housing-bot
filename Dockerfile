# syntax=docker/dockerfile:1

FROM golang:1.18-alpine

WORKDIR /app

COPY *.go ./

RUN go build -o /housing-bot

EXPOSE 80

CMD [ "/housing-bot" ]

