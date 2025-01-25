FROM golang:1.23.5-bookworm AS builder

WORKDIR /app

COPY go.mod go.sum ./
RUN go mod download

COPY . .

RUN go build -o main .

FROM gcr.io/distroless/static-debian12

COPY --from=builder /app/main /app/main

CMD ["./main"]