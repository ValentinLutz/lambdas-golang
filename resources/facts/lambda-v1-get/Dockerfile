FROM golang:1.21.3-alpine3.18 AS builder
WORKDIR /build
COPY ./go.mod ./go.sum ./
RUN go mod download
COPY ./libraries ./libraries
COPY ./resources/facts/lambda-shared ./resources/facts/lambda-shared
COPY ./resources/facts/lambda-v1-get ./resources/facts/lambda-v1-get
RUN go build -tags lambda.norpc -o main ./resources/facts/lambda-v1-get

FROM alpine:3.18
COPY --from=builder /build/main ./main
CMD ["./main"]