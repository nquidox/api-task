FROM golang:1.23.4-alpine3.21 AS builder

WORKDIR /build

COPY go.* ./

RUN go mod download

COPY . .

RUN CGO_ENABLED=0 GOOS=linux go build -ldflags="-s -w" -o app  ./cmd

FROM alpine:3.21

COPY --from=builder /build/app /app

EXPOSE 9000

CMD ["/app"]
