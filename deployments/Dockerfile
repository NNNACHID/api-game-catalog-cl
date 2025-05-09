FROM golang:1.24-alpine AS builder

RUN apk add --no-cache git gcc musl-dev
WORKDIR /app
COPY go.mod go.sum ./
RUN go mod download
COPY . .
RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o catalog-service ./cmd/catalog-service/

FROM alpine:3.21

RUN apk --no-cache add ca-certificates tzdata
WORKDIR /app
COPY --from=builder /app/catalog-service .
COPY --from=builder /app/config ./config
EXPOSE 8080

CMD ["./catalog-service"]
