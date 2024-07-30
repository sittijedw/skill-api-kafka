FROM golang:alpine AS builder
WORKDIR /go/src
COPY . .
RUN go build -o consumer

FROM scratch
COPY --from=builder /go/src/consumer .
ENTRYPOINT ["./consumer"]