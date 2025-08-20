FROM golang:1.24-alpine AS builder

WORKDIR /app

COPY go.mod go.sum ./

RUN go mod download

COPY . .

RUN CGO_ENABLED=0 GOOS=linux go build -ldflags="-w -s" -o main ./cmd/main.go

FROM alpine:3.22

WORKDIR /root/

COPY --from=builder /app/main .

EXPOSE 3001

CMD ["./main"]