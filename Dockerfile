FROM golang:1.23.5 AS builder

WORKDIR /app

COPY . .

RUN go mod tidy

RUN GOOS=linux GOARCH=amd64 go build -o capital-gains ./cmd/main.go

FROM alpine:latest

WORKDIR /root/

COPY --from=builder /app/capital-gains .

CMD ["./capital-gains"]
