FROM golang:1.21.3-alpine3.18 AS builder
WORKDIR /build
COPY ./go.mod ./go.sum ./
RUN go mod download
COPY ./libraries ./libraries
COPY ./lambdas/facts/app-v1-get ./
RUN go build -tags lambda.norpc -o main

FROM alpine:3.18
COPY --from=builder /build/main ./main
CMD ["./main"]