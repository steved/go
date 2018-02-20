FROM golang as builder
WORKDIR /go/src/github.com/steved/go
COPY . /go/src/github.com/steved/go
RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o go .

FROM scratch
COPY --from=builder /go/src/github.com/steved/go/go /
CMD ["/go", "--data=/data"]
