FROM golang:latest

LABEL maintainer="anshita"

WORKDIR /app

COPY go.mod .

COPY go.sum .

RUN go mod download

COPY . .

ENV ADDR 4321

EXPOSE 4321

RUN go build

CMD ["./DOG-API"]