FROM golang:1.23-alpine AS builder

WORKDIR /app

COPY go.mod ./
RUN go mod download

COPY . .

RUN CGO_ENABLED=0 GOOS=linux go build -o wosp-collection .

FROM alpine:3.18

WORKDIR /app

COPY --from=builder /app/wosp-collection .

EXPOSE 8080

CMD ["./wosp-collection"]
