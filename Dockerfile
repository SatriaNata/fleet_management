FROM golang:1.25-alpine

WORKDIR /internal

COPY go.mod go.sum ./
RUN go mod tidy

COPY . .

RUN go build -o main ./cmd/api/main.go

EXPOSE 8080

CMD ["./main"]