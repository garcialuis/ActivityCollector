FROM golang:1.14-alpine

RUN apk update && apk upgrade && apk add --no-cache bash git

RUN go get ./...

ENV SOURCES /home/luis/Documents/Projects/GoLang/ActivityCollector/
COPY . ${SOURCES}

RUN cd ${SOURCES} && CGO_ENABLED=0 go build

ENV BROKER_ADDR amqp://guest:guest@localhost:5672/

WORKDIR ${SOURCES}
CMD ${SOURCES}ActivityCollector