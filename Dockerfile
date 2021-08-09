FROM golang:1.16-alpine

MAINTAINER Tony C. Batista

ENV TIMEZONE America/Sao_Paulo

WORKDIR /app

COPY . .

RUN apk add build-base

RUN go mod download
RUN go mod vendor