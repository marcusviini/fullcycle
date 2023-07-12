FROM golang:latest

WORKDIR /usr/app
COPY . .

RUN go mod tidy
RUN go run main.go