FROM golang:1.14-alpine AS builder
WORKDIR /go/src/app
COPY . .

RUN go get -d -v ./...
RUN go install -v ./...
RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o app .

FROM golang:1.14-alpine
WORKDIR /go/src/app/
COPY --from=builder /go/src/app .
CMD ["./app"]