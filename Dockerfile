FROM golang:1.19 AS builder
WORKDIR /apps
COPY . .
RUN go build -o main main.go

FROM alpine:3.19
WORKDIR /app
COPY --from=builder /app/main .
COPY app.env .

EXPOSE 8080
CMD ["./main"]
ENTRYPOINT [ "/app/start.sh" ]