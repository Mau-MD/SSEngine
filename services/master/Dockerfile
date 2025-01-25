FROM golang:1.23.5-bookworm AS builder

WORKDIR /app

# Copy go mod files first for better caching
COPY ./services/master/go.mod ./services/master/go.sum ./

# Copy the replace script
COPY ./services/master/scripts/replace_go_mod.sh ./scripts/replace_go_mod.sh
# Copy the library
COPY ./libs/proto/ ./output/libs/proto/

RUN chmod +x ./scripts/replace_go_mod.sh
RUN ./scripts/replace_go_mod.sh

RUN go mod download

# Copy the service code (excluding go.mod and go.sum)
COPY ./services/master/*.go ./
RUN CGO_ENABLED=0 GOOS=linux go build -o /master .

# # Use distroless as minimal base image
FROM gcr.io/distroless/static-debian12:latest

WORKDIR /

COPY --from=builder /master /master

EXPOSE 50051

ENTRYPOINT ["/master"]