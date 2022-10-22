FROM golang:1.19.2-alpine3.16

RUN apk update -qq && apk add git

RUN mkdir /app

#ENV GOPATH /app
#ENV GOROOT /app

#ADD . /go/src/app/

WORKDIR /app

#COPY go.mod go.sum ./
#RUN go mod download && go mod verify

## Add this go mod download command to pull in any dependencies
#RUN go mod download
#ENTRYPOINT /app


ADD . /app

COPY . /app

#COPY go.mod ./
#COPY go.sum ./

#RUN go test -v ./

#RUN go mod init github.com/emanuelinacio/go-lang-api
#RUN go mod tidy
#RUN go get github.com/go-redis/redis/v8 

#RUN go get github.com/go-redis/redis/v8
## Our project will now successfully build with the necessary go libraries included

#RUN go build -o server
RUN go test -v ./event/entity
#RUN go run main.go 

#CMD ["/server"]
#CMD /bin/bash
## Our start command which kicks off
## our newly created binary executable