FROM golang:1.16-alpine

RUN apk update && apk add git

RUN mkdir /go/src/app
WORKDIR /go/src/app

COPY . .

RUN go get ./ && go get github.com/pilu/fresh
RUN go build

CMD go run main.go