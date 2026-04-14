FROM golang:1.25-alpine AS builder
WORKDIR /src

COPY go.mod go.sum ./
RUN apk add --no-cache ca-certificates
RUN go mod download

COPY . .
RUN CGO_ENABLED=0 GOOS=linux go build \
    -ldflags="-s -w" \
    -o /src/app ./cmd/main.go

FROM alpine:3.21
RUN apk add --no-cache ca-certificates

WORKDIR /app
COPY --from=builder /src/app /app/app

ENV PORT=8080
EXPOSE 8080

USER 1000
CMD ["/app/app"]