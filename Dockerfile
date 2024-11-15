FROM golang:1.23 AS builder
WORKDIR /app
COPY . .
RUN go build -o main main.go

FROM alpine:3.19
WORKDIR /app
COPY --from=builder /app/main .
COPY start.sh .

EXPOSE 8080
ENTRYPOINT [ "/app/start.sh" ]