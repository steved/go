FROM golang:1.9-alpine

EXPOSE 8067

RUN mkdir /data

WORKDIR /go/src/github.com/steved/go
COPY . .

RUN go build -o bin/go .

CMD ["bin/go", "--data=/data"]
