FROM golang:alpine AS builder
WORKDIR /go/src
COPY . .
RUN go build -o consumer1
RUN go build -o consumer2
RUN go build -o consumer3

FROM scratch
COPY --from=builder /go/src/consumer1 .
COPY --from=builder /go/src/consumer2 .
COPY --from=builder /go/src/consumer3 .
ENTRYPOINT ["./consumer1", "./consumer2", "./consumer3"]