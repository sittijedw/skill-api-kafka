FROM golang:alpine AS builder
WORKDIR /go/src
COPY . .
RUN go build -o api

FROM alpine
COPY --from=builder /go/src/api .
ENTRYPOINT ["./api"]