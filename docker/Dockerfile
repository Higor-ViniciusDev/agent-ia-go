FROM golang:1.24 AS builder

WORKDIR /app

COPY go.mod ./

RUN go mod download

COPY . .

RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 \
    go build -ldflags="-s -w" -o agent-ia ./cmd/api

FROM alpine:3.21

RUN apk --no-cache add ca-certificates curl

WORKDIR /app

COPY --from=builder /app/agent-ia .

EXPOSE 8000

CMD ["./agent-ia"]