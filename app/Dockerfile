# syntax=docker/dockerfile:1

FROM golang:1.20-alpine
ENV GOPROXY="https://goproxy.io"
ENV GO111MODULE="on"
WORKDIR /app
COPY go.mod ./
COPY go.sum ./
RUN go mod download
COPY . .

RUN go build -o /depay ./app/main.go

EXPOSE 5009

CMD [ "/depay" ]

