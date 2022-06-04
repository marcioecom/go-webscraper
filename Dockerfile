FROM golang:1.18-alpine

WORKDIR /app

RUN apk update && \
    apk add libc-dev && \
    apk add gcc && \
    apk add make

COPY . . 

RUN go mod download && go mod tidy && go mod verify

EXPOSE 3000

RUN go build -o bin/webscrapper main.go

ADD https://raw.githubusercontent.com/eficode/wait-for/v2.1.0/wait-for /usr/local/bin/wait-for

RUN chmod +rx /usr/local/bin/wait-for entrypoint.sh

ENTRYPOINT [ "sh", "entrypoint.sh" ]
