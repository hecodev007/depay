# syntax=docker/dockerfile:1

FROM golang:1.20-alpine
ENV GOPROXY="https://goproxy.io"
ENV GO111MODULE="on"
WORKDIR /admin
COPY go.mod ./
COPY go.sum ./
RUN go mod download
COPY . .
RUN go build -o /adminapp ./admin/main.go

EXPOSE 5010

CMD [ "/adminapp" ]

