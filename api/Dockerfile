FROM golang:alpine AS builder
WORKDIR /go/src
COPY . .
RUN go build -o api

FROM scratch
COPY --from=builder /go/src/api .
ENTRYPOINT ["./api"]