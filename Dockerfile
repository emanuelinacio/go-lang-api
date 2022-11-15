FROM golang:latest

RUN mkdir -p /app
WORKDIR /app
ADD . /app

#RUN go test -v ./event/entity

RUN go build -o /main

EXPOSE 8080

CMD [ "/main" ]