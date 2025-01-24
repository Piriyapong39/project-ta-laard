FROM golang:1.22-alpine

WORKDIR /app/user-service

COPY user-service/go.mod user-service/go.sum ./

RUN go mod tidy

COPY user-service/ .

RUN go build -o main ./cmd/main.go

EXPOSE 8080

CMD ["./main"]