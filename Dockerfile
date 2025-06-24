FROM golang:latest AS builder

WORKDIR /app
COPY . .

RUN go mod download

RUN make build

FROM alpine

WORKDIR /app
COPY --from=builder /app/HTTP_IO_BOUND .

EXPOSE 8080

CMD ["./HTTP_IO_BOUND"]