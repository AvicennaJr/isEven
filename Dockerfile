FROM golang:1.20-alpine

RUN apk update && apk upgrade && \
    apk add --no-cache bash git openssh

WORKDIR /app

COPY . .

RUN go build -o example ./example/main.go

CMD ["./example"]
