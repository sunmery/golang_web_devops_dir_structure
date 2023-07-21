FROM golang

RUN mkdir /app

ADD . /app

WORKDIR /app

RUN go build -o main ./cmd/server/main.go

EXPOSE 4000

CMD ["./app/main"]
