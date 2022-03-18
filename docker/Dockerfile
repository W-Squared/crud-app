FROM golang:1.17

RUN mkdir /app

ADD . /app

WORKDIR /app

RUN go get -d -v .
RUN go install -v .
RUN go build -o main .

CMD ["/app/main"]

EXPOSE 8080