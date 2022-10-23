FROM golang:1.19.2-alpine3.16

RUN apk update -qq && apk add git

RUN mkdir /app

WORKDIR /app

#RUN go mod download && go mod verify

#RUN go mod download

ADD . /app

#RUN go mod init github.com/emanuelinacio/go-lang-api
#RUN go mod tidy
#RUN go get github.com/go-redis/redis/v8 

RUN go test -v ./event/entity
RUN go run main.go 

#CMD ["/server"]
#CMD /bin/bash
## Our start command which kicks off
## our newly created binary executable