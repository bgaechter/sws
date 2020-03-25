FROM golang:1.14-buster AS builder
WORKDIR /go/src/app
COPY . .

RUN go get -d -v ./...
RUN go install -v ./...
RUN go build -o app .

FROM gcr.io/distroless/cc-debian10
COPY --from=builder /go/src/app /
CMD ["/app"]