FROM golang:1.20 AS builder

WORKDIR /app

COPY src/ .

RUN go mod init capital-gains && go mod tidy

RUN go build -o capital-gains .

FROM scratch:latest

WORKDIR /root/

COPY --from=builder /app/capital-gains .

CMD ["./capital-gains"]