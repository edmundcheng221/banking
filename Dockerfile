FROM golang:1.19-alpine3.18 as builder

WORKDIR /app
COPY . .

RUN go build -o main main.go
FROM alpine:3.18
WORKDIR /app
COPY --from=builder /app/main .

EXPOSE 8080

CMD ["/app/main"]
